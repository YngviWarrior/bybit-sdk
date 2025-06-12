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

func (s *bybit) LiveOrder(stopChan <-chan struct{}) {
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

	// var ConnID string
	var Authenticated bool
	var Subscribed bool
	go func(w http.ResponseWriter, r *http.Request) {
		for {
			var response bybitstructs.OldWebSocketAuthResponse
			var responseData bybitstructs.LiveOrderData
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Fatal("LOV5:", err)
			}
			Log(string(msg))
			if err = json.Unmarshal(msg, &response); err != nil {
				log.Panic("LOV5 00")
			}

			if response.RetMsg != "" {
				log.Println("Erro LOV5 02:", response.RetMsg)
				return
			}

			if Subscribed {
				if err = json.Unmarshal(msg, &responseData); err != nil {
					log.Panic("LOV5 03")
				}

				if responseData.RetCode == 0 {
					data, err := json.Marshal(responseData.Data)
					if err != nil {
						log.Panic("LOV5 03.1 ", err)
					}

					mqConn.Publish("order", "direct", responseData.Topic, data)
				} else {
					log.Panic("LOV5 05: ", err)
				}
			}

			if Authenticated {
				err = json.Unmarshal(msg, &response)
				if err != nil {
					log.Println("Subscription Failed", err)
				}

				Subscribed = true
			}

			// ConnID = response.ConnID
			Authenticated = true
		}
	}(nil, nil)

	message, err := json.Marshal(authMessage)
	if err != nil {
		log.Println("LOV5 02:", err)
		return
	}

	// Enviar uma mensagem para o servidor WebSocket
	err = conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Fatal("LOV5 03:", err)
	}

	subscription := fmt.Sprintf(`{"op":"subscribe","args":["%s"]}`, `order`)

	// Enviar uma mensagem para o servidor WebSocket
	err = conn.WriteMessage(websocket.TextMessage, []byte(subscription))
	if err != nil {
		log.Fatal("Erro ao enviar mensagem:", err)
	}

	// Enviar um heart beat (ping) a cada 20 segundos (opcional)
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := conn.WriteMessage(websocket.PingMessage, []byte(`{
				"op": "ping"
			}`))
			if err != nil {
				log.Fatal("LOV5 04:", err)
			}

		case <-stopChan:
			fmt.Println("Encerrando conexão...")
			// Envia uma mensagem de encerramento antes de fechar
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Encerrando conexão"))
			if err != nil {
				log.Println("LOV5 05:", err)
				return
			}
			// Dá tempo para o servidor processar a mensagem de fechamento
			time.Sleep(1 * time.Second)
			return
		}
	}
}
