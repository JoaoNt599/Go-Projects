package rabbitmq

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

// ConnectRabbitMQ connects to RabbitMQ and returns a channel
func ConnectRabbitMQ() (*amqp.Connection, *amqp.Channel) {
	rabbitURL := "amqp://guest:guest@rabbitmq:5672/"

	// Attempting to connect to RabbitMQ up to 5 times
	var conn *amqp.Connection
	var err error
	for i := 1; i <= 5; i++ {
		conn, err = amqp.Dial(rabbitURL)
		if err == nil {
			break
		}
		fmt.Printf("Attempt %d: Waiting for RabbitMQ...\n", i)
		time.Sleep(3 * time.Second)
	}
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		os.Exit(1)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
	}

	_, err = ch.QueueDeclare(
		"chat_messages",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	fmt.Println("Connected to RabbitMQ")
	return conn, ch
}
