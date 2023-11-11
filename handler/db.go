package handler

type NFTListDate struct {
	CreateTime    int64          `json:"create_time"`
	DenomId       string         `json:"denom_id"`
	DenomName     string         `json:"denom_name"`
	LastBlockTime int64          `json:"last_block_time"`
	NftId         string         `json:"nft_id"`
	NftName       string         `json:"nft_name"`
	Owner         string         `json:"owner"`
	TokenData     string         `json:"tokenData"`
	TokenUri      string         `json:"tokenUri"`
	TransferData  []TransferData `json:"transfer_data"`
}

type TransferData struct {
	TransferTime string      `json:"Time"`
	BlockHeight  int64       `json:"blockHeight"`
	DenomId      string      `json:"denomId"`
	Fee          interface{} `json:"fee"`
	Sender       string      `json:"sender"`
	Status       int         `json:"status"`
	TxHash       string      `json:"txHash"`
	TxType       []string    `json:"txType"`
	TxId         int64       `json:"tx_id"`
}
