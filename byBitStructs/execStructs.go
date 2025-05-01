package bybitstructs

type ExecutionMessage struct {
	Topic        string      `json:"topic"`
	ID           string      `json:"id"`
	CreationTime uint64      `json:"creationTime"`
	Data         []Execution `json:"data"`
}

type Execution struct {
	Category        string `json:"category"`
	Symbol          string `json:"symbol"`
	ClosedSize      string `json:"closedSize"`
	ExecFee         string `json:"execFee"`
	ExecId          string `json:"execId"`
	ExecPrice       string `json:"execPrice"`
	ExecQty         string `json:"execQty"`
	ExecType        string `json:"execType"`
	ExecValue       string `json:"execValue"`
	FeeRate         string `json:"feeRate"`
	TradeIv         string `json:"tradeIv"`
	MarkIv          string `json:"markIv"`
	BlockTradeId    string `json:"blockTradeId"`
	MarkPrice       string `json:"markPrice"`
	IndexPrice      string `json:"indexPrice"`
	UnderlyingPrice string `json:"underlyingPrice"`
	LeavesQty       string `json:"leavesQty"`
	OrderId         string `json:"orderId"`
	OrderLinkId     string `json:"orderLinkId"`
	OrderPrice      string `json:"orderPrice"`
	OrderQty        string `json:"orderQty"`
	OrderType       string `json:"orderType"`
	StopOrderType   string `json:"stopOrderType"`
	Side            string `json:"side"`
	ExecTime        string `json:"execTime"`
	IsLeverage      string `json:"isLeverage"`
	IsMaker         bool   `json:"isMaker"`
	Seq             uint64 `json:"seq"`
	MarketUnit      string `json:"marketUnit"`
}
