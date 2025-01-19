package bybitService_test

import (
	"os"
	"testing"
	"time"

	service "github.com/YngviWarrior/BybitSDK/bybitService"
)

var bybit service.BybitServiceInterface

func TestMain(m *testing.M) {
	bybit = service.NewBybitService(os.Getenv("BYBIT_PUBLIC_KEY"), os.Getenv("BYBIT_SECRET_KEY"))
}

func TestLivePublic(t *testing.T) {
	stopChan := make(chan struct{})
	go bybit.LivePublic("kline.1.BTCUSDT", stopChan)

	time.Sleep(time.Second * 5)
	stopChan <- struct{}{}
}

func TestGetWalletInfo(t *testing.T) {
	response := bybit.GetWalletInfo()

	if response.RetCode != 0 {
		t.Fatalf(response.RetMsg)
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

	// t.Logf("Server TimeSecond: %v", response)
}
