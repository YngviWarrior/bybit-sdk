package bybitSDK_test

import (
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	service "github.com/YngviWarrior/BybitSDK"
	bybitstructs "github.com/YngviWarrior/BybitSDK/byBitStructs"
	"github.com/joho/godotenv"
)

var bybit service.BybitServiceInterface

func TestMain(m *testing.M) {
	err := godotenv.Load(`../.env`)

	if err != nil {
		log.Fatal(".env file is missing")
	}

	bybit = service.NewBybitService(os.Getenv("BYBIT_PUBLIC_KEY"), os.Getenv("BYBIT_SECRET_KEY"))

	code := m.Run()
	os.Exit(code)
}

func TestLivePublic(t *testing.T) {
	stopChan := make(chan struct{})
	go bybit.LivePublic("kline.1.BTCUSDT", stopChan)

	time.Sleep(time.Second * 5)
	stopChan <- struct{}{}
}

func TestLiveExec(t *testing.T) {
	order := &bybitstructs.CreateTradeParams{
		// ReqID: "Test-003",
		Header: struct {
			XBAPITimestamp  string "json:\"X-BAPI-TIMESTAMP\""
			XBAPIRecvWindow string "json:\"X-BAPI-RECV-WINDOW\""
		}{
			XBAPITimestamp:  strconv.FormatInt(time.Now().Unix()*1000, 10),
			XBAPIRecvWindow: "10000",
		},
		Op: "order.create",
		Args: []struct {
			Symbol      string "json:\"symbol\""
			Side        string "json:\"side\""
			OrderType   string "json:\"orderType\""
			Qty         string "json:\"qty\""
			Category    string "json:\"category\""
			TimeInForce string "json:\"timeInForce\""
		}{
			{
				Symbol:      "BTCUSDT",
				Side:        "Buy",
				OrderType:   "Market",
				Qty:         "10",
				Category:    "spot",
				TimeInForce: "GTC",
			},
		},
	}

	createOrderChan := make(chan *bybitstructs.CreateTradeParams)
	cancelOrderChan := make(chan *bybitstructs.CancelTradeParams)
	stopChan := make(chan struct{})

	go bybit.LiveExec(createOrderChan, cancelOrderChan, stopChan)
	time.Sleep(time.Second * 3)

	createOrderChan <- order

	time.Sleep(time.Second * 5)
	stopChan <- struct{}{}
}

func TestLiveOrder(t *testing.T) {
	stopChan := make(chan struct{})
	go bybit.LiveOrders(stopChan)

	time.Sleep(time.Second * 5)
	stopChan <- struct{}{}
}

func TestGetWalletInfo(t *testing.T) {
	response := bybit.GetWalletInfo()

	if response.RetCode != 0 {
		t.Fatal(response.RetMsg)
	}

	for _, v := range response.Account.Balance {
		if v.WalletBalance != "0" {
			t.Logf("Coin: %v, Balance: %v , Transferable: %v", v.Coin, v.WalletBalance, v.TransferBalance)
		}
	}
}

func TestGetServerTimestamp(t *testing.T) {
	response := bybit.GetServerTimestamp()

	if response == 0 {
		t.Fatalf("Server TimeSecond is 0")
	}
}

func TestOrderHistory(t *testing.T) {
	response := bybit.OrderHistory(&bybitstructs.OrderHistoryParams{
		Category: "spot",
		Symbol:   "BTCUSDT",
	})

	if response.RetCode != 0 {
		t.Fatal(response.RetMsg)
	}

	if len(response.Result.List) == 0 {
		t.Fatalf("Order history list is empty")
	}

}

func TestOpenOrders(t *testing.T) {
	response := bybit.OpenOrders(&bybitstructs.OpenOrderParams{
		Category: "spot",
		Symbol:   "BTCUSDT",
	})

	if response.RetCode != 0 {
		t.Fatal(response.RetMsg)
	}
}

func TestCancelOrder(t *testing.T) {
	response := bybit.CancelOrders(&bybitstructs.CancelOrderParams{})

	if response.RetCode != 0 {
		t.Fatal(response.RetMsg)
	}
}
