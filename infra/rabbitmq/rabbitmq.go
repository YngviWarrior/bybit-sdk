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

	return &rabbitmq{
		Conn: conn,
	}
}
