package bybitSDK

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	bybitstructs "github.com/YngviWarrior/BybitSDK/byBitStructs"
)

func (s *bybit) setHeaders(req *http.Request, signature string, timestamp int64) {
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	req.Header.Add("X-BAPI-SIGN", signature)
	req.Header.Add("X-BAPI-API-KEY", os.Getenv("BYBIT_API_KEY"))
	req.Header.Add("X-BAPI-TIMESTAMP", fmt.Sprintf("%v", timestamp))
	req.Header.Add("X-BAPI-RECV-WINDOW", recvWindow)
}

func (s *bybit) gerSign(timestamp int64, queryStr string) (sign string) {
	ruleStr := fmt.Sprintf("%v%s%s", timestamp, os.Getenv("BYBIT_API_KEY"), recvWindow)
	sign = ruleStr + queryStr
	return
}

func (s *bybit) gerQueryString(q url.Values, values map[string]string) (queryStr string) {
	for k, v := range values {
		q.Add(k, v)
	}

	queryStr = q.Encode()
	return
}

func (s *bybit) GetServerTimestamp() (timestamp int64) {
	s.setUrl()
	client := &http.Client{}

	req, err := http.NewRequest("GET", BASE_URL+"/v5/market/time", nil)

	if err != nil {
		log.Println("BBST 01: ", err)
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		log.Println("BBST exec: ", err)
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var response bybitstructs.GetServerTimestamp
	err = json.Unmarshal(bodyBytes, &response)

	if err != nil {
		log.Println("BBST 02: ", err)
	}

	timestamp, _ = strconv.ParseInt(response.Result.TimeNano, 10, 64)

	defer resp.Body.Close()

	return timestamp / 1_000_000
}

func (s *bybit) GetKlines(params *bybitstructs.GetKlinesParams) (response bybitstructs.GetKlinesResponse) {
	s.setUrl()
	client := &http.Client{}

	req, err := http.NewRequest("GET", BASE_URL+"/v5/market/kline", nil)

	if err != nil {
		log.Println("BBGK 01: ", err)
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	timeResp := s.GetServerTimestamp()

	req.URL.RawQuery = s.gerQueryString(req.URL.Query(), map[string]string{
		"category": "spot",
		"interval": params.Interval,
		"symbol":   params.Symbol,
		"start":    fmt.Sprintf("%v", params.StartTime),
		"end":      fmt.Sprintf("%v", params.EndTime),
		"limit":    fmt.Sprintf("%v", params.Limit),
	})

	sig := hmac.New(sha256.New, []byte(os.Getenv("BYBIT_SECRET_KEY")))
	sig.Write([]byte(s.gerSign(timeResp, req.URL.RawQuery)))

	s.setHeaders(req, hex.EncodeToString(sig.Sum(nil)), timeResp)
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Req klines exec: ", err)
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(bodyBytes, &response)

	if err != nil {
		log.Println("BBGK 02: ", err)
	}

	defer resp.Body.Close()

	return
}

func (s *bybit) CancelOrders(params *bybitstructs.CancelOrderParams) (response bybitstructs.CancelOrderResponse) {
	s.setUrl()
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", BASE_URL+"/v5/order/cancel", nil)

	if err != nil {
		log.Println("BBCO 01: ", err.Error())
		return
	}

	timeResp := s.GetServerTimestamp()
	q := req.URL.Query()

	s.gerQueryString(q, map[string]string{
		"orderId": params.OrderId,
	})

	req.URL.RawQuery = q.Encode()

	s.gerSign(timeResp, req.URL.RawQuery)
	sig := hmac.New(sha256.New, []byte(os.Getenv("BYBIT_SECRET_KEY")))
	sig.Write([]byte(s.gerSign(timeResp, req.URL.RawQuery)))

	s.setHeaders(req, hex.EncodeToString(sig.Sum(nil)), timeResp)
	resp, err := client.Do(req)

	if err != nil {
		log.Println("BBCO 02: ", err.Error())
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("BBCO 03: ", err.Error())
		return
	}

	err = json.Unmarshal(bodyBytes, &response)

	if err != nil {
		log.Println("BBCO 04: ", err.Error())
		return
	}

	defer resp.Body.Close()

	return
}

func (s *bybit) OpenOrders(params *bybitstructs.OpenOrderParams) (response bybitstructs.OpenOrderResponse) {
	s.setUrl()
	client := &http.Client{}
	req, err := http.NewRequest("GET", BASE_URL+"/v5/order/realtime", nil)

	if err != nil {
		log.Println("BBOO 01: ", err.Error())
		return
	}

	timeResp := s.GetServerTimestamp()
	q := req.URL.Query()

	s.gerQueryString(q, map[string]string{
		"symbol":   params.Symbol,
		"category": params.Category,
		"orderId":  params.OrderId,
	})

	req.URL.RawQuery = q.Encode()

	s.gerSign(timeResp, req.URL.RawQuery)
	sig := hmac.New(sha256.New, []byte(os.Getenv("BYBIT_SECRET_KEY")))
	sig.Write([]byte(s.gerSign(timeResp, req.URL.RawQuery)))

	s.setHeaders(req, hex.EncodeToString(sig.Sum(nil)), timeResp)
	resp, err := client.Do(req)

	if err != nil {
		log.Println("BBOO 02: ", err.Error())
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("BBOO 03: ", err.Error())
		return
	}

	if bodyBytes != nil {
		err = json.Unmarshal(bodyBytes, &response)

		if err != nil {
			log.Println("BBOO 04: ", err.Error())
			return
		}
	}

	defer resp.Body.Close()

	return
}

func (s *bybit) OrderHistory(params *bybitstructs.OrderHistoryParams) (response bybitstructs.OrderHistoryResponse) {
	s.setUrl()
	client := &http.Client{}
	req, err := http.NewRequest("GET", BASE_URL+"/v5/order/history", nil)

	if err != nil {
		log.Println("BBOH 01: ", err.Error())
		return
	}

	timeResp := s.GetServerTimestamp()
	q := req.URL.Query()

	s.gerQueryString(q, map[string]string{
		"category": params.Category,
		"symbol":   params.Symbol,
		"orderId":  params.OrderId,
	})

	req.URL.RawQuery = q.Encode()

	s.gerSign(timeResp, req.URL.RawQuery)
	sig := hmac.New(sha256.New, []byte(os.Getenv("BYBIT_SECRET_KEY")))
	sig.Write([]byte(s.gerSign(timeResp, req.URL.RawQuery)))

	s.setHeaders(req, hex.EncodeToString(sig.Sum(nil)), timeResp)
	resp, err := client.Do(req)

	if err != nil {
		log.Println("BBOH 02: ", err.Error())
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("BBOH 03: ", err.Error())
		return
	}

	err = json.Unmarshal(bodyBytes, &response)

	if bodyBytes != nil {
		if err != nil {
			log.Println("BBOH 04: ", err.Error())
			return
		}
	}

	defer resp.Body.Close()

	return
}

func (s *bybit) CreateOrder(params *bybitstructs.OrderParams) (response *bybitstructs.OrderResponse) {
	s.setUrl()
	client := &http.Client{}
	req, err := http.NewRequest("POST", BASE_URL+"/v5/order/create", nil)

	if err != nil {
		log.Println("BBC 01: ", err.Error())
		return
	}

	timeResp := s.GetServerTimestamp()

	q := req.URL.Query()

	s.gerQueryString(q, map[string]string{
		"symbol":      params.Symbol,
		"qty":         params.OrderQty,
		"side":        params.Side,
		"type":        params.OrderType,
		"timeInForce": params.TimeInForce,
		"price":       params.OrderPrice,
	})

	req.URL.RawQuery = q.Encode()

	s.gerSign(timeResp, req.URL.RawQuery)
	sig := hmac.New(sha256.New, []byte(os.Getenv("BYBIT_SECRET_KEY")))
	sig.Write([]byte(s.gerSign(timeResp, req.URL.RawQuery)))

	s.setHeaders(req, hex.EncodeToString(sig.Sum(nil)), timeResp)
	resp, err := client.Do(req)

	if err != nil {
		log.Println("BBC 02: ", err.Error())
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("BBC 03: ", err.Error())
		return
	}

	err = json.Unmarshal(bodyBytes, &response)

	if err != nil {
		log.Println("BBC 04: ", err.Error())
		return
	}

	defer resp.Body.Close()

	return
}

func (s *bybit) GetWalletInfo() (response *bybitstructs.GetWalletInfoResponse) {
	s.setUrl()
	client := &http.Client{}
	req, err := http.NewRequest("GET", BASE_URL+"/v5/asset/transfer/query-account-coins-balance", nil)

	if err != nil {
		log.Println("BBGWI 01: ", err.Error())
		return
	}

	timeResp := s.GetServerTimestamp()

	req.URL.RawQuery = s.gerQueryString(req.URL.Query(), map[string]string{
		"accountType": "UNIFIED",
		"coin":        "BTC, ETH, USDT",
	})

	sig := hmac.New(sha256.New, []byte(os.Getenv("BYBIT_SECRET_KEY")))
	sig.Write([]byte(s.gerSign(timeResp, req.URL.RawQuery)))

	s.setHeaders(req, hex.EncodeToString(sig.Sum(nil)), timeResp)
	resp, err := client.Do(req)

	if err != nil {
		log.Println("BBGWI 02: ", err.Error())
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("BBGWI 03: ", err.Error())
		return
	}

	err = json.Unmarshal(bodyBytes, &response)

	if err != nil {
		log.Println("BBGWI 04: ", err.Error())
		return
	}

	defer resp.Body.Close()

	return
}
