package bybitSDK

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	bybitstructs "github.com/YngviWarrior/bybit-sdk/byBitStructs"
	"github.com/gorilla/websocket"
)

func (s *bybit) LivePublic(topic []string, stopChan <-chan struct{}) {
	s.setUrl()
	var topics string
	for _, v := range topic {
		fmt.Println(v)
		topics += fmt.Sprintf(`"%s",`, v)
	}
	topics = topics[:len(topics)-1]
	fmt.Println(topics)

	conn, _, err := websocket.DefaultDialer.Dial(BASE_URL_WSS+"/v5/public/spot", nil)
	if err != nil {
		log.Fatal("Erro ao conectar ao WebSocket:", err)
	}
	defer conn.Close()

	var ConnID string
	go func() {
		var Subscribed bool
		for {
			var responseSubscription bybitstructs.WebSocketAuthResponse
			var responseKline bybitstructs.SocketKlineResponse
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Fatal("Erro ao ler mensagem:", err)
			}

			if err = json.Unmarshal(msg, &responseSubscription); err != nil {
				log.Panic("LPV5 00")
			}
			fmt.Println(responseSubscription)
			if Subscribed {
				if err = json.Unmarshal(msg, &responseKline); err != nil {
					log.Panic("LPV5 01")
				}
				fmt.Println(responseKline)
				data, err := json.Marshal(responseKline.Data)
				if err != nil {
					log.Panic("LPV5 02 ", err)
				}

				if err := s.Conn.Set(context.Background(), responseKline.Topic, data, 0).Err(); err != nil {
					log.Panic("LPV5 03 ", err)
				}
			}

			Subscribed = true
			ConnID = responseSubscription.ConnID
		}
	}()

	fmt.Println("Conectado ao WebSocket:", BASE_URL_WSS+"/v5/public/spot")
	subscription := fmt.Sprintf(`{"op":"subscribe","args":[%s]}`, topics)
	fmt.Println(subscription)
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
				"success": true,
				"ret_msg": "pong",
				"conn_id": "`+ConnID+`",
				"op": "ping"
			}`))
			if err != nil {
				log.Fatal("Erro ao enviar ping:", err)
			}
		case <-stopChan:
			fmt.Println("Encerrando conexão...")
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Encerrando conexão"))
			if err != nil {
				log.Println("Erro ao enviar mensagem de encerramento:", err)
				return
			}
			// Dá tempo para o servidor processar a mensagem de fechamento
			time.Sleep(1 * time.Second)
			return
		}
	}
}
