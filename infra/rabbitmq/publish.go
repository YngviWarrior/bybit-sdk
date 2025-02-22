package rabbitmq

import (
	"log"

	amqp091 "github.com/rabbitmq/amqp091-go"
)

func (r *rabbitmq) Publish(exchangeName string, exchangeType string, queueName string, data []byte) {
	ch, err := r.Conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	defer ch.Close()

	if exchangeName != "" {
		err = ch.ExchangeDeclare(
			exchangeName, // Nome do Exchange
			exchangeType, // Tipo Fanout
			false,        // Não persistente
			false,        // Não autodelete
			false,
			false,
			nil,
		)
		if err != nil {
			log.Fatal("Erro ao declarar exchange:", err)
		}
	}

	_, err = ch.QueueDeclare(
		queueName,
		false, // Durable: false (não persiste após reiniciar o RabbitMQ)
		false, // Auto Delete: false (não apaga quando ninguém estiver consumindo)
		false, // Exclusive: false (pode ser acessada por múltiplos consumidores)
		false, // No Wait: false
		nil,   // Arguments
	)

	if err != nil {
		log.Fatal("RP 00: ", err)
	}

	err = ch.Publish(
		exchangeName,
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
