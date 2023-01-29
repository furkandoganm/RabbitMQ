package consumers

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"projects/RabbitMQWithMultipleChannel_3/publishers"
	"strconv"
)

// func ConsumerRMQ(ch *amqp.Channel, cVal *chan int) {
func ConsumerRMQ(ch *amqp.Channel, q string) {
	defer ch.Close()

	metaDatas, err := ch.Consume(
		q,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("error consuming val in %s: %v", q, err)
	}
	fmt.Println(q)
	//d := <-metaDatas

	//val, _ := strconv.Atoi(string(d.Body))
	//val += 10
	//if q == "Queue1" {
	//	publishers.PublisherRMQ(ch, "Queue2", val)
	//} else if q == "Queue2" {
	//	publishers.PublisherRMQ(ch, "Queue3", val)
	//}
	//forever := make(chan int)

	for d := range metaDatas {
		val, _ := strconv.Atoi(string(d.Body))
		val += 10
		if q == "Queue1" {
			publishers.PublisherRMQ(ch, "Queue2", val)
		} else if q == "Queue2" {
			publishers.PublisherRMQ(ch, "Queue3", val)
		}
	}

	//<-forever
}
