package bybitSDK

import (
	"os"

	bybitstructs "github.com/YngviWarrior/bybit-sdk/byBitStructs"
	"github.com/YngviWarrior/bybit-sdk/infra/rabbitmq"
)

var BASE_URL = "https://api.bybit.com"
var BASE_URL_WSS = "wss://stream.bybit.com"

const recvWindow = "10000"

type BybitServiceInterface interface {
	LivePublic(topic []string, stopChan <-chan struct{})
	LiveOrder(stopChan <-chan struct{})
	LiveExec(stopChan <-chan struct{})
	LivePosition(stopChan <-chan struct{})
	LiveTrade(order <-chan *bybitstructs.OrderRequest, stopChan <-chan struct{})

	GetServerTimestamp() (response int64)
	GetWalletInfo() (response *bybitstructs.GetWalletInfoResponse)
	CreateOrder(params *bybitstructs.OrderParams) (response *bybitstructs.OrderResponse)
	OrderHistory(params *bybitstructs.OrderHistoryParams) (response bybitstructs.OrderHistoryResponse)
	OpenOrders(params *bybitstructs.OpenOrderParams) (response bybitstructs.OpenOrderResponse)
	CancelOrders(params *bybitstructs.CancelOrderParams) (response bybitstructs.CancelOrderResponse)
	GetKlines(params *bybitstructs.GetKlinesParams) (response bybitstructs.GetKlinesResponse)
}

type bybit struct {
	Conn rabbitmq.RabbitMQInterface
}

func (s *bybit) setUrl() {
	switch os.Getenv("ENVIROMENT") {
	case "testnet":
		BASE_URL = "https://api-testnet.bybit.com"
		BASE_URL_WSS = "wss://stream-testnet.bybit.com"
	default:
		BASE_URL = "https://api.bybit.com"
		BASE_URL_WSS = "wss://stream.bybit.com"
	}
}

func NewBybitService(publicKey, secretKey string) BybitServiceInterface {
	messagebroker := rabbitmq.NewRabbitMQConnection()

	return &bybit{
		Conn: messagebroker,
	}
}
