package bybitstructs

type OrderMessage struct {
	Topic        string        `json:"topic"`
	ID           string        `json:"id"`
	CreationTime uint64        `json:"creationTime"`
	Data         []OrderDetail `json:"data"`
}

type OrderDetail struct {
	Category              string `json:"category"`
	Symbol                string `json:"symbol"`
	OrderID               string `json:"orderId"`
	OrderLinkID           string `json:"orderLinkId"`
	BlockTradeID          string `json:"blockTradeId"`
	Side                  string `json:"side"`
	PositionIdx           uint64 `json:"positionIdx"`
	OrderStatus           string `json:"orderStatus"`
	CancelType            string `json:"cancelType"`
	RejectReason          string `json:"rejectReason"`
	TimeInForce           string `json:"timeInForce"`
	IsLeverage            string `json:"isLeverage"`
	Price                 string `json:"price"`
	Qty                   string `json:"qty"`
	AvgPrice              string `json:"avgPrice"`
	LeavesQty             string `json:"leavesQty"`
	LeavesValue           string `json:"leavesValue"`
	CumExecQty            string `json:"cumExecQty"`
	CumExecValue          string `json:"cumExecValue"`
	CumExecFee            string `json:"cumExecFee"`
	OrderType             string `json:"orderType"`
	StopOrderType         string `json:"stopOrderType"`
	OrderIv               string `json:"orderIv"`
	TriggerPrice          string `json:"triggerPrice"`
	TakeProfit            string `json:"takeProfit"`
	StopLoss              string `json:"stopLoss"`
	TriggerBy             string `json:"triggerBy"`
	TpTriggerBy           string `json:"tpTriggerBy"`
	SlTriggerBy           string `json:"slTriggerBy"`
	TriggerDirection      uint64 `json:"triggerDirection"`
	PlaceType             string `json:"placeType"`
	LastPriceOnCreated    string `json:"lastPriceOnCreated"`
	CloseOnTrigger        bool   `json:"closeOnTrigger"`
	ReduceOnly            bool   `json:"reduceOnly"`
	SmpGroup              uint64 `json:"smpGroup"`
	SmpType               string `json:"smpType"`
	SmpOrderID            string `json:"smpOrderId"`
	SlLimitPrice          string `json:"slLimitPrice"`
	TpLimitPrice          string `json:"tpLimitPrice"`
	MarketUnit            string `json:"marketUnit"`
	CreatedTime           string `json:"createdTime"`
	UpdatedTime           string `json:"updatedTime"`
	FeeCurrency           string `json:"feeCurrency"`
	SlippageTolerance     string `json:"slippageTolerance"`
	SlippageToleranceType string `json:"slippageToleranceType"`
}
