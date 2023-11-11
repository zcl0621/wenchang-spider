package handler

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
	"wenchang-spider/utils"
)

// ...

func ExportToCsv(nftList []NFTListDate, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	// 写入表头
	headers := []string{
		"Create Time",
		"Denom ID",
		"Denom Name",
		"Last Block Time",
		"NFT ID",
		"NFT Name",
		"Owner",
		"Token Data",
		"Token Uri",
		"Transfer Time",
		"Block Height",
		"Denom ID",
		"Fee",
		"Sender",
		"Status",
		"Tx Hash",
		"Tx Type",
		"Tx ID",
	}
	writer.Write(headers)

	// 写入数据行
	for _, nft := range nftList {
		for _, transfer := range nft.TransferData {
			row := []string{
				strconv.FormatInt(nft.CreateTime, 10),
				nft.DenomId,
				nft.DenomName,
				strconv.FormatInt(nft.LastBlockTime, 10),
				nft.NftId,
				nft.NftName,
				nft.Owner,
				nft.TokenData,
				nft.TokenUri,
				transfer.TransferTime,
				strconv.FormatInt(transfer.BlockHeight, 10),
				transfer.DenomId,
				utils.ConvertInterfaceToString(transfer.Fee),
				transfer.Sender,
				strconv.Itoa(transfer.Status),
				transfer.TxHash,
				strings.Join(transfer.TxType, ","),
				strconv.FormatInt(transfer.TxId, 10),
			}
			writer.Write(row)
		}
	}

	writer.Flush()
	return nil
}
