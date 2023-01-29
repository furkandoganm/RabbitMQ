package main

import (
	"log"
	"projects/RabbitMQWithMultipleChannel_3/configs"
	"projects/RabbitMQWithMultipleChannel_3/consumers"
	"projects/RabbitMQWithMultipleChannel_3/models"
	"projects/RabbitMQWithMultipleChannel_3/publishers"
	"time"
)

func main() {

	amqpC := configs.ConnectionRMQ()
	defer amqpC.Close()

	ch, err := amqpC.Channel()
	if err != nil {
		log.Fatalln("error creating channel: ", err)
	}
	//defer ch.Close()
	for k, _ := range models.QueueList {
		_, err := ch.QueueDeclare(
			k,
			false,
			false,
			false,
			false, nil,
		)
		if err != nil {
			log.Fatalf("error declaring queue (%s): %v", k, err)
		}
		//if k == "Queue1" {
		//	publishers.PublisherRMQ(ch, k, 0)
		//}
	}

	go func() {
		for k, _ := range models.QueueList {
			if k != "Queue3" {
				go consumers.ConsumerRMQ(ch, k)
			}
		}
	}()

	for {
		publishers.PublisherRMQ(ch, "Queue1", 0)
		time.Sleep(time.Second * 1)
	}

}
