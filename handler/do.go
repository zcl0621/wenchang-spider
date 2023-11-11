package handler

import (
	"fmt"
	"github.com/go-rod/rod"
	"time"
	"wenchang-spider/browser"
)

func Do(nftName string) {
	d, p := HandlerList(nftName)
	if d == nil || p == nil {
		panic("没有获取到任何数据")
	}
	for i := range d {
		HandlerDetail(p, &d[i])
	}
	fmt.Printf("全部信息已获取完成，正在导出。。。")
	e := ExportToCsv(d, nftName+".csv")
	if e != nil {
		panic(e)
	}
}

func HandlerList(nftName string) ([]NFTListDate, *rod.Page) {
	p := browser.LiveBrowser.MustPage("https://tianhe.wenchang.bianjie.ai/#/mto/nfts")
	p.MustElement("#app > div.page_container.content")
	time.Sleep(time.Second * 1)
	p.MustEval(`(nftName)=>{
                     document.querySelector('#app > div.page_container.content > section > div.list_container_header > div.list_container_extra > div > div.search_filter_input_content > div > input').value = nftName;
               }`, nftName)
	p.MustEval(`(nftName)=> {
            document.querySelector('#app').__vue__.$children[4].$children[0].$children[1]._data.nftId= nftName;
            document.querySelector('#app').__vue__.$children[4].$children[0].$children[1]._data.query= nftName;
        }`, nftName)
	p.MustEval(`()=>{
                         document.querySelector('#app > div.page_container.content > section > div.list_container_header > div.list_container_extra > div > div.button_container > div.search_filter_refresh_content > button').click();
                   }`)
	var AllNftList []NFTListDate
	time.Sleep(1 * time.Second)
	nftCount := p.MustEval(`()=> {
                    return document.querySelector('#app').__vue__.$children[4].$children[0].$children[1].nftCount
		}`).Int()
	getCount := 0
	fmt.Printf("正在获取列表数据 共%d条 \n", nftCount)
	for {
		date := p.MustEval(`() =>{
                    return document.querySelector('#app').__vue__.$children[4].$children[0].$children[1].pageData
            }`)
		var nftList []NFTListDate
		e := date.Unmarshal(&nftList)
		if e != nil {
			panic(e)
		}
		AllNftList = append(AllNftList, nftList...)
		getCount += len(nftList)
		fmt.Printf("正在获取列表数据 共%d条 已获取%d条 \n", nftCount, getCount)
		if getCount >= nftCount {
			break
		}
		p.MustEval(`()=> {
			document.querySelector("#app > div.page_container.content > section > div.list_container_footer > div.list_container_pagination > div > button.next_button").click();
		}`)
		time.Sleep(time.Second * 1)

	}
	return AllNftList, p
}

func HandlerDetail(p *rod.Page, oneNft *NFTListDate) {
	p.MustEval(`
		(denom_id,nft_id) => {
			let url = "/mto/nfts/detail?denom="+denom_id+"&tokenId="+nft_id;
			document.querySelector('#app').__vue__.$router.push(url);
			}
	`, oneNft.DenomId, oneNft.NftId)
	p.MustElement("#app > div.page_container.nft_token_container.content")
	time.Sleep(time.Second * 1)
	date := p.MustEval(`()=> {
		 return document.querySelector('#app').__vue__.$children[4].$children[0].$children[2].listData
	}`)
	var tData []TransferData
	e := date.Unmarshal(&tData)
	if e != nil {
		panic(e)
	}
	oneNft.TransferData = tData
	fmt.Printf("获取到详情数据 块为: %s\n", oneNft.NftId)
}
