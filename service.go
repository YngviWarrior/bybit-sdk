package bybitSDK

import (
	"context"
	"log"
	"os"

	bybitstructs "github.com/YngviWarrior/bybit-sdk/byBitStructs"
	"github.com/redis/go-redis/v9"
)

var BASE_URL = "https://api.bybit.com"
var BASE_URL_WSS = "wss://stream.bybit.com"

const recvWindow = "10000"

type BybitServiceInterface interface {
	LivePublic(topic string, stopChan <-chan struct{})
	LiveOrders(stopChan <-chan struct{})
	LiveExec(createOrderChan <-chan *bybitstructs.CreateTradeParams, cancelOrderChan <-chan *bybitstructs.CancelTradeParams, stopChan <-chan struct{})

	GetServerTimestamp() (response int64)
	GetWalletInfo() (response *bybitstructs.GetWalletInfoResponse)
	CreateOrder(params *bybitstructs.OrderParams) (response *bybitstructs.OrderResponse)
	OrderHistory(params *bybitstructs.OrderHistoryParams) (response bybitstructs.OrderHistoryResponse)
	OpenOrders(params *bybitstructs.OpenOrderParams) (response bybitstructs.OpenOrderResponse)
	CancelOrders(params *bybitstructs.CancelOrderParams) (response bybitstructs.CancelOrderResponse)
}

type bybit struct {
	Conn *redis.Client
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
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Endereço do servidor Redis
		Password: "",               // Senha (se houver)
		DB:       0,                // Número do banco de dados
	})

	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Erro ao conectar ao Redis: %v", err)
	}

	log.Println("Conexão com Redis bem-sucedida!")

	return &bybit{
		Conn: rdb,
	}
}
