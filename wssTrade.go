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

func (s *bybit) LiveTrade(order <-chan *bybitstructs.OrderRequest, stopChan <-chan struct{}) {
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
	var Subscribed bool
	go func(w http.ResponseWriter, r *http.Request) {
		for {
			var response bybitstructs.WebSocketAuthResponse
			var responseData bybitstructs.OrderResponse
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Fatal("Erro to read message:", err)
			}

			err = json.Unmarshal(msg, &response)
			if err != nil {
				log.Println("Erro to unmarshal message:", err)
			}

			if response.RetMsg != "" {
				log.Println("Erro LTV5 00:", response.RetMsg)
				return
			}

			fmt.Printf("WSST: %v\n", string(msg))
			if Subscribed {
				err = json.Unmarshal(msg, &responseData)
				if err != nil {
					log.Println("Erro to unmarshal message:", err)
				}

				if responseData.RetCode == 0 {
					data, err := json.Marshal(responseData)
					if err != nil {
						log.Panic("LTV5 01 ", err)
					}

					mqConn.Publish("livetrade", "direct", responseData.Op, data)
				} else {
					log.Panic("LTV5 05: ", err)
				}
			}

			if Authenticated {
				err = json.Unmarshal(msg, &response)
				if err != nil {
					log.Println("Subscription Failed", err)
				}
			}

			Subscribed = true
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
		log.Fatal("Erro to send message:", err)
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
				log.Fatal("Erro to send ping:", err)
			}
		case orderMessage := <-order:
			bytes, err := json.Marshal(orderMessage)
			if err != nil {
				log.Println("Erro to marshal order message:", err)
				return
			}
			fmt.Println("Order message: ", string(bytes))
			err = conn.WriteMessage(websocket.TextMessage, bytes)
			if err != nil {
				log.Fatal("Erro to send ping:", err)
			}
		case <-stopChan:
			fmt.Println("Encerrando conexão...")
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Encerrando conexão"))
			if err != nil {
				log.Println("Erro to send closure message:", err)
				return
			}

			time.Sleep(1 * time.Second)
			return
		}
	}
}
