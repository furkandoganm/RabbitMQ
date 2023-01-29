package publishers

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"strconv"
	"time"
)

// func PublisherRMQ(ch *amqp.Channel, q string, cVal chan int) {
// func PublisherRMQ(ch *amqp.Channel, q string, cVal *chan int) {
func PublisherRMQ(ch *amqp.Channel, q string, val int) {
	//var val int = <-*cVal
	//var val int = <-cVal

	if val == 0 {
		val = Rand()
	}
	//defer ch.Close()

	err := ch.Publish(
		"",
		q,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(strconv.Itoa(val)),
		})
	fmt.Printf("%s kuyruğuna %v değeri eklendi \n", q, val)
	if err != nil {
		log.Fatalf("error publishing data to %s: %v", q, err)
	}
	//time.Sleep(time.Second * 1)

}

func Rand() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(100)
}
