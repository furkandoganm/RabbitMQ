package configs

import (
	"github.com/streadway/amqp"
	"log"
)

func ConnectionRMQ() *amqp.Connection {
	conn, err := amqp.Dial(EnvRabbitMQ())
	if err != nil {
		log.Fatalln("error RabbitMQ connection: ", err)
	}
	return conn
}
