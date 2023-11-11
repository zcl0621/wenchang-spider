package main

import (
	"fmt"
	"os"
	"wenchang-spider/browser"
	"wenchang-spider/handler"
)

func main() {
	if len(os.Args) > 1 {
		arg1 := os.Args[1]
		fmt.Println("初始化环境中。。。")
		browser.InitLiveBrowser()
		fmt.Println("开始获取数据。。。")
		handler.Do(arg1)
		fmt.Println("任务已完成")
	} else {
		fmt.Println("请输入需要获取的NFT项目名称，请输入全名，否则查询出的结果会存在误差")
	}
}
