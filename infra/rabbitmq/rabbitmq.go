package rabbitmq

import (
	"log"
	"os"

	amqp091 "github.com/rabbitmq/amqp091-go"
)

type rabbitmq struct {
	Conn *amqp091.Connection
}

type RabbitMQInterface interface {
	Publish(exchangeName string, exchangeType string, queueName string, data []byte)
}

func NewRabbitMQConnection() RabbitMQInterface {
	conn, err := amqp091.Dial(os.Getenv("RABBIT_MQ"))
	if err != nil {
		log.Fatal(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	defer ch.Close()

	err = ch.ExchangeDeclare(
		"order",  // name
		"direct", // type (ou "topic", "fanout", etc.)
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare exchange: %v", err)
	}

	err = ch.ExchangeDeclare(
		"execution", // name
		"direct",    // type (ou "topic", "fanout", etc.)
		true,        // durable
		false,       // auto-deleted
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare exchange: %v", err)
	}

	err = ch.ExchangeDeclare(
		"livetrade", // name
		"direct",    // type (ou "topic", "fanout", etc.)
		true,        // durable
		false,       // auto-deleted
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare exchange: %v", err)
	}

	err = ch.ExchangeDeclare(
		"klines", // name
		"fanout", // type (ou "topic", "fanout", etc.)
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare exchange: %v", err)
	}

	return &rabbitmq{
		Conn: conn,
	}
}
