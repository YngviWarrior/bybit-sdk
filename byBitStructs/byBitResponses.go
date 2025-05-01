package bybitstructs

type SocketKlineResponse struct {
	Type  string `json:"type"`
	Topic string `json:"topic"`
	Data  []struct {
		Start     int64  `json:"start"`
		End       int64  `json:"end"`
		Interval  string `json:"interval"`
		Open      string `json:"open"`
		Close     string `json:"close"`
		High      string `json:"high"`
		Low       string `json:"low"`
		Volume    string `json:"volume"`
		Turnover  string `json:"turnover"`
		Confirm   bool   `json:"confirm"`
		Timestamp int64  `json:"timestamp"`
	} `json:"data"`
	Ts int64 `json:"ts"`
}

// Result encapsula os resultados principais
type GetWalletInfoResponse struct {
	RetCode    int               `json:"retCode"`
	RetExtInfo map[string]string `json:"retExtInfo"`
	RetMsg     string            `json:"retMsg"`
	Time       float64           `json:"time"`
	Account    struct {
		AccountType string `json:"accountType"`
		Balance     []struct {
			Bonus           string `json:"bonus"`
			Coin            string `json:"coin"`
			TransferBalance string `json:"transferBalance"`
			WalletBalance   string `json:"walletBalance"`
		} `json:"balance"`
		MemberID string `json:"memberId"`
	} `json:"result"`
}

type GetServerTimestamp struct {
	RetCode    int               `json:"retCode"`
	RetExtInfo map[string]string `json:"retExtInfo"`
	RetMsg     string            `json:"retMsg"`
	Time       float64           `json:"time"`
	Result     struct {
		TimeSecond string `json:"timeSecond"`
		TimeNano   string `json:"timeNano"`
	} `json:"result"`
}

type GetKlinesResponse struct {
	RetCode    int               `json:"retCode"`
	RetExtInfo map[string]string `json:"retExtInfo"`
	RetMsg     string            `json:"retMsg"`
	Time       float64           `json:"time"`
	Category   string            `json:"category"`
	Result     struct {
		Symbol   string     `json:"symbol"`
		Category string     `json:"category"`
		List     [][]string `json:"list"`
	} `json:"result"`
}

type CancelOrderResponse struct {
	RetCode int64  `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
	ExtCode string `json:"ext_code"`
	ExtInfo string `json:"ext_info"`
	Result  struct {
		AccountId    string `json:"accountId"`
		Symbol       string `json:"symbol"`
		OrderLinkId  string `json:"orderLinkId"`
		OrderId      string `json:"orderId"`
		TransactTime string `json:"transactTime"`
		Price        string `json:"price"`
		OrigQty      string `json:"origQty"`
		ExecutedQty  string `json:"executedQty"`
		Status       string `json:"status"`
		TimeInForce  string `json:"timeInForce"`
		Type         string `json:"type"`
		Side         string `json:"side"`
	} `json:"result"`
}

type OpenOrderResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		List []struct {
			OrderID            string `json:"orderId"`
			OrderLinkID        string `json:"orderLinkId"`
			BlockTradeID       string `json:"blockTradeId"`
			Symbol             string `json:"symbol"`
			Price              string `json:"price"`
			Qty                string `json:"qty"`
			Side               string `json:"side"`
			IsLeverage         string `json:"isLeverage"`
			PositionIdx        int    `json:"positionIdx"`
			OrderStatus        string `json:"orderStatus"`
			CancelType         string `json:"cancelType"`
			RejectReason       string `json:"rejectReason"`
			AvgPrice           string `json:"avgPrice"`
			LeavesQty          string `json:"leavesQty"`
			LeavesValue        string `json:"leavesValue"`
			CumExecQty         string `json:"cumExecQty"`
			CumExecValue       string `json:"cumExecValue"`
			CumExecFee         string `json:"cumExecFee"`
			TimeInForce        string `json:"timeInForce"`
			OrderType          string `json:"orderType"`
			StopOrderType      string `json:"stopOrderType"`
			OrderIv            string `json:"orderIv"`
			TriggerPrice       string `json:"triggerPrice"`
			TakeProfit         string `json:"takeProfit"`
			StopLoss           string `json:"stopLoss"`
			TpTriggerBy        string `json:"tpTriggerBy"`
			SlTriggerBy        string `json:"slTriggerBy"`
			TriggerDirection   int    `json:"triggerDirection"`
			TriggerBy          string `json:"triggerBy"`
			LastPriceOnCreated string `json:"lastPriceOnCreated"`
			ReduceOnly         bool   `json:"reduceOnly"`
			CloseOnTrigger     bool   `json:"closeOnTrigger"`
			SmpType            string `json:"smpType"`
			SmpGroup           int    `json:"smpGroup"`
			SmpOrderID         string `json:"smpOrderId"`
			TpslMode           string `json:"tpslMode"`
			TpLimitPrice       string `json:"tpLimitPrice"`
			SlLimitPrice       string `json:"slLimitPrice"`
			PlaceType          string `json:"placeType"`
			CreatedTime        string `json:"createdTime"`
			UpdatedTime        string `json:"updatedTime"`
		} `json:"list"`
		NextPageCursor string `json:"nextPageCursor"`
		Category       string `json:"category"`
	} `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

type OrderHistoryResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		List []struct {
			OrderID            string `json:"orderId"`
			OrderLinkID        string `json:"orderLinkId"`
			BlockTradeID       string `json:"blockTradeId"`
			Symbol             string `json:"symbol"`
			Price              string `json:"price"`
			Qty                string `json:"qty"`
			Side               string `json:"side"`
			IsLeverage         string `json:"isLeverage"`
			PositionIdx        int    `json:"positionIdx"`
			OrderStatus        string `json:"orderStatus"`
			CancelType         string `json:"cancelType"`
			RejectReason       string `json:"rejectReason"`
			AvgPrice           string `json:"avgPrice"`
			LeavesQty          string `json:"leavesQty"`
			LeavesValue        string `json:"leavesValue"`
			CumExecQty         string `json:"cumExecQty"`
			CumExecValue       string `json:"cumExecValue"`
			CumExecFee         string `json:"cumExecFee"`
			TimeInForce        string `json:"timeInForce"`
			OrderType          string `json:"orderType"`
			StopOrderType      string `json:"stopOrderType"`
			OrderIv            string `json:"orderIv"`
			TriggerPrice       string `json:"triggerPrice"`
			TakeProfit         string `json:"takeProfit"`
			StopLoss           string `json:"stopLoss"`
			TpTriggerBy        string `json:"tpTriggerBy"`
			SlTriggerBy        string `json:"slTriggerBy"`
			TriggerDirection   int    `json:"triggerDirection"`
			TriggerBy          string `json:"triggerBy"`
			LastPriceOnCreated string `json:"lastPriceOnCreated"`
			ReduceOnly         bool   `json:"reduceOnly"`
			CloseOnTrigger     bool   `json:"closeOnTrigger"`
			SmpType            string `json:"smpType"`
			SmpGroup           int    `json:"smpGroup"`
			SmpOrderID         string `json:"smpOrderId"`
			TpslMode           string `json:"tpslMode"`
			TpLimitPrice       string `json:"tpLimitPrice"`
			SlLimitPrice       string `json:"slLimitPrice"`
			PlaceType          string `json:"placeType"`
			CreatedTime        string `json:"createdTime"`
			UpdatedTime        string `json:"updatedTime"`
		} `json:"list"`
		NextPageCursor string `json:"nextPageCursor"`
		Category       string `json:"category"`
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type AllCoinsResponse struct {
	RetCode int64  `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
	ExtCode string `json:"ext_code"`
	ExtInfo string `json:"ext_info"`
	Result  struct {
		Balances []struct {
			Coin     string `json:"coin"`
			CoinId   string `json:"coinId"`
			CoinName string `json:"coinName"`
			Free     string `json:"free"`
			Locked   string `json:"locked"`
			Total    string `json:"total"`
		} `json:"balances"`
	} `json:"result"`
	TimeNow          string `json:"time_now"`
	RateLimitStatus  int64  `json:"rate_limit_status"`
	RateLimitResetMs int64  `json:"rate_limit_reset_ms"`
	RateLimit        int64  `json:"rate_limit"`
}

type WebSocketAuthResponse struct {
	Success bool   `json:"success"`
	RetMsg  string `json:"ret_msg"`
	Op      string `json:"op"`
	ConnID  string `json:"conn_id"`
}

type OldWebSocketAuthResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Op      string `json:"op"`
	ConnID  string `json:"connId"`
}

type LiveOrderData struct {
	OldWebSocketAuthResponse
	ID           string `json:"id"`
	Topic        string `json:"topic"`
	CreationTime int64  `json:"creationTime"`
	Data         []struct {
		Symbol             string `json:"symbol"`
		OrderID            string `json:"orderId"`
		Side               string `json:"side"`
		OrderType          string `json:"orderType"`
		CancelType         string `json:"cancelType"`
		Price              string `json:"price"`
		Qty                string `json:"qty"`
		OrderIv            string `json:"orderIv"`
		TimeInForce        string `json:"timeInForce"`
		OrderStatus        string `json:"orderStatus"`
		OrderLinkID        string `json:"orderLinkId"`
		LastPriceOnCreated string `json:"lastPriceOnCreated"`
		ReduceOnly         bool   `json:"reduceOnly"`
		LeavesQty          string `json:"leavesQty"`
		LeavesValue        string `json:"leavesValue"`
		CumExecQty         string `json:"cumExecQty"`
		CumExecValue       string `json:"cumExecValue"`
		AvgPrice           string `json:"avgPrice"`
		BlockTradeID       string `json:"blockTradeId"`
		PositionIdx        int    `json:"positionIdx"`
		CumExecFee         string `json:"cumExecFee"`
		ClosedPnl          string `json:"closedPnl"`
		CreatedTime        string `json:"createdTime"`
		UpdatedTime        string `json:"updatedTime"`
		RejectReason       string `json:"rejectReason"`
		StopOrderType      string `json:"stopOrderType"`
		TpslMode           string `json:"tpslMode"`
		TriggerPrice       string `json:"triggerPrice"`
		TakeProfit         string `json:"takeProfit"`
		StopLoss           string `json:"stopLoss"`
		TpTriggerBy        string `json:"tpTriggerBy"`
		SlTriggerBy        string `json:"slTriggerBy"`
		TpLimitPrice       string `json:"tpLimitPrice"`
		SlLimitPrice       string `json:"slLimitPrice"`
		TriggerDirection   int    `json:"triggerDirection"`
		TriggerBy          string `json:"triggerBy"`
		CloseOnTrigger     bool   `json:"closeOnTrigger"`
		Category           string `json:"category"`
		PlaceType          string `json:"placeType"`
		SmpType            string `json:"smpType"`
		SmpGroup           int    `json:"smpGroup"`
		SmpOrderID         string `json:"smpOrderId"`
		FeeCurrency        string `json:"feeCurrency"`
	} `json:"data"`
}

type LiveExecResponse struct {
	OldWebSocketAuthResponse
	Data struct {
		OrderID     string `json:"orderId"`
		OrderLinkID string `json:"orderLinkId"`
	} `json:"data"`
	Header struct {
		XBapiLimit               string `json:"X-Bapi-Limit"`
		XBapiLimitStatus         string `json:"X-Bapi-Limit-Status"`
		XBapiLimitResetTimestamp string `json:"X-Bapi-Limit-Reset-Timestamp"`
		TraceID                  string `json:"Traceid"`
		TimeNow                  string `json:"Timenow"`
	} `json:"header"`
}
