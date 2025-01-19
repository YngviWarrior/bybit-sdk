package service_test

import (
	"os"
	"testing"
	"time"

	bybit "github.com/YngviWarrior/BybitSDK/bybit"
)

var service bybit.BybitServiceInterface

func TestMain(m *testing.M) {
	service = bybit.NewBybitService(os.Getenv("BYBIT_PUBLIC_KEY"), os.Getenv("BYBIT_SECRET_KEY"))
}

func TestLivePublic(t *testing.T) {
	stopChan := make(chan struct{})
	service.LivePublicV5("kline.1.BTCUSDT", stopChan)

	time.Sleep(time.Second * 5)
	stopChan <- struct{}{}
}
