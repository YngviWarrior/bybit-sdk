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

func (s *bybit) LiveExec(stopChan <-chan struct{}) {
	s.setUrl()
	mqConn := rabbitmq.NewRabbitMQConnection()

	conn, _, err := websocket.DefaultDialer.Dial(BASE_URL_WSS+"/v5/private", nil)
	if err != nil {
		log.Fatal("Erro ao conectar ao WebSocket:", err)
	}
	defer conn.Close()

	fmt.Println("Conectado ao WebSocket:", BASE_URL_WSS+"/v5/private")

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
			var responseData bybitstructs.ExecutionMessage
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Fatal("Erro to read message:", err)
			}

			err = json.Unmarshal(msg, &response)
			if err != nil {
				log.Println("Erro to unmarshal message:", err)
			}

			if response.RetMsg != "" {
				log.Println("Erro LEV5 00:", response.RetMsg)
				return
			}
			fmt.Println("response: ", string(msg))
			if Subscribed {
				err = json.Unmarshal(msg, &responseData)
				if err != nil {
					log.Println("Erro to unmarshal message:", err)
				}

				if responseData.Topic == "execution" {
					data, err := json.Marshal(responseData.Data)
					if err != nil {
						log.Panic("LEV5 01 ", err)
					}

					mqConn.Publish("", "", responseData.Topic, data)
				} else {
					log.Panic("LEV5 05: ", err)
				}
			}

			if Authenticated {
				err = json.Unmarshal(msg, &response)
				if err != nil {
					log.Println("Subscription Failed", err)
				}

				Subscribed = true
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
		log.Fatal("Erro to sent message:", err)
	}

	subscription := fmt.Sprintf(`{"op":"subscribe","args":["%s"]}`, `execution`)

	err = conn.WriteMessage(websocket.TextMessage, []byte(subscription))
	if err != nil {
		log.Fatal("Erro to sent message:", err)
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
		case <-stopChan:
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Encerrando conexão"))
			if err != nil {
				log.Println("Erro to send close message:", err)
				return
			}

			time.Sleep(1 * time.Second)
			return
		}
	}
}
