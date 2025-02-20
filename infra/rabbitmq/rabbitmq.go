package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type rabbitmq struct {
	Conn *amqp.Connection
}

type RabbitMQInterface interface {
	Publish(string, []byte)
}

func NewRabbitMQConnection() RabbitMQInterface {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Conexão com RabbitMQ bem-sucedida!")

	return &rabbitmq{
		Conn: conn,
	}
}
