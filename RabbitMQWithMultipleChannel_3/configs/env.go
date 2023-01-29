package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvRabbitMQ() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error .env RabbitMQ: ", err)
	}

	rabbitMQConfig := os.Getenv("RabbitMQ")
	return rabbitMQConfig
}
