package server

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func StartServer() *amqp.Connection {
	fmt.Println("RabbitMQ in Golang: Getting started tutorial")

	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Printf("failed to connect to RabbitMQ: %v", err)
		return nil
	}
	fmt.Println("Successfully connected to RabbitMQ instance")
	return connection
}
