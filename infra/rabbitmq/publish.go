package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

func (r *rabbitmq) Publish(queueName string, data []byte) {
	ch, err := r.Conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	defer ch.Close()

	queue, err := ch.QueueDeclare(
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
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		},
	)

	if err != nil {
		log.Fatal("RP 01: ", err)
	}

}
