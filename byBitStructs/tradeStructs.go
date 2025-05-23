package bybitstructs

type OrderRequest struct {
	ReqID  string          `json:"reqId"`
	Header RequestHeader   `json:"header"`
	Op     string          `json:"op"`
	Args   []OrderArgument `json:"args"`
}

type RequestHeader struct {
	Timestamp  string `json:"X-BAPI-TIMESTAMP"`
	RecvWindow string `json:"X-BAPI-RECV-WINDOW"`
	Referer    string `json:"Referer"`
}

type OrderArgument struct {
	Symbol      string `json:"symbol"`
	Side        string `json:"side"`
	OrderType   string `json:"orderType"`
	Qty         string `json:"qty"`
	Price       string `json:"price,omitempty"`
	Category    string `json:"category"`
	TimeInForce string `json:"timeInForce"`
	OrderLinkId string `json:"orderLinkId"`
}

/* -------- */

type OrderResponse struct {
	ReqID      string              `json:"reqId"`
	RetCode    int                 `json:"retCode"`
	RetMsg     string              `json:"retMsg"`
	Op         string              `json:"op"`
	Data       OrderResponseData   `json:"data"`
	RetExtInfo map[string]any      `json:"retExtInfo"`
	Header     OrderResponseHeader `json:"header"`
	ConnID     string              `json:"connId"`
}

type OrderResponseData struct {
	OrderID     string `json:"orderId"`
	OrderLinkID string `json:"orderLinkId"`
}

type OrderResponseHeader struct {
	LimitResetTimestamp string `json:"X-Bapi-Limit-Reset-Timestamp"`
	TraceID             string `json:"Traceid"`
	TimeNow             string `json:"Timenow"`
	Limit               string `json:"X-Bapi-Limit"`
	LimitStatus         string `json:"X-Bapi-Limit-Status"`
}
