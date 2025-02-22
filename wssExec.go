package bybitSDK

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	bybitstructs "github.com/YngviWarrior/bybit-sdk/byBitStructs"
	"github.com/YngviWarrior/bybit-sdk/infra/rabbitmq"
	"github.com/gorilla/websocket"
)

func (s *bybit) LiveExec(createOrderChan <-chan *bybitstructs.CreateTradeParams, cancelOrderChan <-chan *bybitstructs.CancelTradeParams, stopChan <-chan struct{}) {
	s.setUrl()
	mqConn := rabbitmq.NewRabbitMQConnection()

	conn, _, err := websocket.DefaultDialer.Dial(BASE_URL_WSS+"/v5/trade", nil)
	if err != nil {
		log.Fatal("Erro ao conectar ao WebSocket:", err)
	}
	defer conn.Close()

	fmt.Println("Conectado ao WebSocket:", BASE_URL_WSS+"/v5/trade")

	expires := time.Now().UnixNano()/1e6 + 10000
	mac := hmac.New(sha256.New, []byte(os.Getenv("BYBIT_SECRET_KEY")))
	mac.Write([]byte(fmt.Sprintf("GET/realtime%d", expires)))
	sign := hex.EncodeToString(mac.Sum(nil))

	authMessage := map[string]interface{}{
		"op":   "auth",
		"args": []string{os.Getenv("BYBIT_API_KEY"), fmt.Sprint(expires), sign},
	}

	var ConnID string
	var Authenticated bool
	go func(w http.ResponseWriter, r *http.Request) {
		for {
			var response bybitstructs.WebSocketAuthResponse
			var responseData bybitstructs.LiveExecResponse
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Fatal("Erro ao ler mensagem:", err)
			}

			err = json.Unmarshal(msg, &response)
			if err != nil {
				log.Println("Erro ao fazer unmarshal da mensagem:", err)
			}

			if response.RetMsg != "" {
				log.Println("Erro LEV5 00:", response.RetMsg)
				return
			}

			// fmt.Printf("Mensagem recebida LEV5: %v\n", string(msg))
			if Authenticated {
				err = json.Unmarshal(msg, &responseData)
				if err != nil {
					log.Println("Erro ao fazer unmarshal da mensagem:", err)
				}

				if responseData.RetCode == 0 {
					data, err := json.Marshal(responseData.Data)
					if err != nil {
						log.Panic("LEV5 01 ", err)
					}

					mqConn.Publish("", "", responseData.Op, data)
				} else {
					log.Panic("LEV5 05: ", err)
				}
			}

			ConnID = response.ConnID
			Authenticated = true
		}
	}(nil, nil)

	message, err := json.Marshal(authMessage)
	if err != nil {
		log.Println("Failed to marshal auth message:", err)
		return
	}
	fmt.Println("auth msg: ", string(message))

	err = conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Fatal("Erro ao enviar mensagem:", err)
	}

	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := conn.WriteMessage(websocket.PingMessage, []byte(`{
				"success": true,
				"ret_msg": "pong",
				"conn_id": "`+ConnID+`",
				"op": "ping"
			}`))
			if err != nil {
				log.Fatal("Erro ao enviar ping:", err)
			}
		case order := <-createOrderChan:
			fmt.Println("Enviando ordem...")

			subscriptionMessage, err := json.Marshal(order)
			if err != nil {
				log.Println("Erro ao fazer marshal da mensagem:", err)
				return
			}

			if err := conn.WriteMessage(websocket.TextMessage, subscriptionMessage); err != nil {
				log.Println("Erro ao enviar mensagem de ordem:", err)
				return
			}

		case order := <-cancelOrderChan:
			fmt.Println("Cancelando ordem...")

			subscriptionMessage, err := json.Marshal(order)
			if err != nil {
				log.Println("Erro ao fazer marshal da mensagem:", err)
				return
			}

			if err := conn.WriteMessage(websocket.TextMessage, subscriptionMessage); err != nil {
				log.Println("Erro ao enviar mensagem de cancelamento:", err)
				return
			}

		case <-stopChan:
			fmt.Println("Encerrando conexão...")
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Encerrando conexão"))
			if err != nil {
				log.Println("Erro ao enviar mensagem de encerramento:", err)
				return
			}

			time.Sleep(1 * time.Second)
			return
		}
	}
}
