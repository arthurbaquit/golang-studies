package publisher

import (
	"context"
	"fmt"

	"github.com/arthurbaquit/golang-studies/rabbitmq/server"
	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishIntoQueue(text, queueKey string) {
	connection := server.StartServer()
	if connection == nil {
		panic("Connection not established")
	}
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	// publishing a message
	err = channel.PublishWithContext(
		context.Background(),
		"",       // exchange
		queueKey, // key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(text),
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully published message")
}

func DeclareQueue(queueKey string) {
	connection := server.StartServer()
	if connection == nil {
		panic("Connection not established")
	}
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	// declaring queue with its properties over the the channel opened
	queue, err := channel.QueueDeclare(
		queueKey, // name
		false,    // durable
		false,    // auto delete
		false,    // exclusive
		false,    // no wait
		nil,      // args
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Queue status:", queue)
}
