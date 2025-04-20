package bybitSDK_test

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	service "github.com/YngviWarrior/bybit-sdk"
	bybitstructs "github.com/YngviWarrior/bybit-sdk/byBitStructs"
	"github.com/joho/godotenv"
)

var bybit service.BybitServiceInterface

func TestMain(m *testing.M) {
	err := godotenv.Load(`.env`)

	if err != nil {
		log.Fatal(".env file is missing")
	}

	bybit = service.NewBybitService(os.Getenv("BYBIT_PUBLIC_KEY"), os.Getenv("BYBIT_SECRET_KEY"))

	code := m.Run()
	os.Exit(code)
}

func TestGetKlines(t *testing.T) {
	response := bybit.GetKlines(&bybitstructs.GetKlinesParams{
		Symbol:   "BTCUSDT",
		Interval: "1",
	})
	fmt.Println(response)
}

func TestCreateOrder(t *testing.T) {
	response := bybit.CreateOrder(&bybitstructs.OrderParams{
		Category:    "spot",
		Symbol:      "BTCUSDT",
		OrderQty:    "10",
		Side:        "Buy",
		OrderType:   "Market",
		TimeInForce: "GTC",
		// OrderPrice:  "100.000",
		// OrderLinkId: "test2",
	})
	fmt.Println(response)
}

func TestLivePublic(t *testing.T) {
	stopChan := make(chan struct{})
	go bybit.LivePublic([]string{"kline.1.BTCUSDT", "kline.1.ETHUSDT"}, stopChan)

	time.Sleep(time.Second * 20)
	stopChan <- struct{}{}
}

func TestLiveOrders(t *testing.T) {
	stopChan := make(chan struct{})
	go bybit.LiveOrders(stopChan)

	time.Sleep(time.Second * 20)
	stopChan <- struct{}{}
}

func TestLiveExec(t *testing.T) {
	stopChan := make(chan struct{})

	go bybit.LiveExec(stopChan)
	time.Sleep(time.Second * 30)

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
	response := bybit.CancelOrders(&bybitstructs.CancelOrderParams{
		Category:    "spot",
		Symbol:      "BTCUSDT",
		OrderId:     "1932902878251685376",
		OrderLinkId: "test",
	})

	if response.RetCode != 0 {
		t.Fatal(response.RetMsg)
	}
}
