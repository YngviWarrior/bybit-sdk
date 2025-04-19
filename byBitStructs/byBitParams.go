package bybitstructs

type GetKlinesParams struct {
	Symbol    string `json:"symbol"`
	Interval  string `json:"interval"`
	Limit     int64  `json:"limit"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
}

type CancelOrderParams struct {
	OrderId     string `json:"orderId"`
	OrderLinkId string `json:"orderLinkId"`
}

type OpenOrderParams struct {
	Category string `json:"category"`
	Symbol   string `json:"symbol"`
	OrderId  string `json:"orderId"`
	// OpenOnly int64   `json:"openOnly"` 0 default
	Limit int64 `json:"limit"`
}

type OrderHistoryParams struct {
	Category  string `json:"category"`
	Symbol    string `json:"symbol"`
	OrderId   string `json:"orderId"`
	Limit     int64  `json:"limit"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
}

type OrderParams struct {
	OrderId       string `json:"orderId"`
	OrderLinkId   string `json:"orderLinkId"`
	Symbol        string `json:"symbol"`
	CreateTime    string `json:"createTime"`
	OrderPrice    string `json:"price"`
	OrderQty      string `json:"qty"`
	OrderType     string `json:"orderType"`
	Side          string `json:"side"`
	Status        string `json:"status"`
	TimeInForce   string `json:"timeInForce"`
	AccountId     string `json:"accountId"`
	OrderCategory int64  `json:"orderCategory"`
	TriggerPrice  string `json:"triggerPrice"`

	Category string `json:"category"`
}

type CreateTradeParams struct {
	ReqID  string `json:"reqId"`
	Header struct {
		XBAPITimestamp  string `json:"X-BAPI-TIMESTAMP"`
		XBAPIRecvWindow string `json:"X-BAPI-RECV-WINDOW"`
	} `json:"header"`
	Op   string `json:"op"`
	Args []struct {
		Symbol      string `json:"symbol"`
		Side        string `json:"side"`
		OrderType   string `json:"orderType"`
		Qty         string `json:"qty"`
		Category    string `json:"category"`
		TimeInForce string `json:"timeInForce"`
	} `json:"args"`
}

type CancelTradeParams struct {
	ReqID  string `json:"reqId"`
	Header struct {
		XBAPITimestamp  string `json:"X-BAPI-TIMESTAMP"`
		XBAPIRecvWindow string `json:"X-BAPI-RECV-WINDOW"`
	} `json:"header"`
	Op   string `json:"op"`
	Args []struct {
		Category    string `json:"category"`
		Symbol      string `json:"symbol"`
		Op          string `json:"op"`
		OrderId     string `json:"orderId"`
		OrderLinkId string `json:"orderLinkId"`
		OrderFilter string `json:"orderFilter"`
	} `json:"args"`
}
