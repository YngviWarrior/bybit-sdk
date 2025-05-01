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
	Price       string `json:"price"`
	Category    string `json:"category"`
	TimeInForce string `json:"timeInForce"`
}
