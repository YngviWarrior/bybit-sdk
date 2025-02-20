package rabbitmq

import (
	"log"

	amqp091 "github.com/rabbitmq/amqp091-go"
)

func (r *rabbitmq) Publish(queueName string, data []byte) {
	ch, err := r.Conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	defer ch.Close()

	err = ch.ExchangeDeclare(
		"klines", // Nome do Exchange
		"fanout", // Tipo Fanout
		false,    // Não persistente
		false,    // Não autodelete
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal("RP 00: ", err)
	}

	defer ch.Close()

	err = ch.Publish(
		"klines",
		"",
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        data,
		},
	)

	if err != nil {
		log.Fatal("RP 01: ", err)
	}

}
