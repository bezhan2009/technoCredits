package brokers

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"technoCredits/internal/app/models"
	"technoCredits/pkg/logger"
)

var RabbitConn *amqp.Connection
var RabbitChannel *amqp.Channel

func ConnectToRabbitMq(params models.RabbitParams) error {
	var err error
	RabbitConn, err = amqp.Dial(params.URLConn)
	if err != nil {
		log.Fatal(err)
	}

	RabbitChannel, err = RabbitConn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func CloseRabbitMQ() error {
	if RabbitChannel != nil {
		err := RabbitChannel.Close()
		if err != nil {
			return err
		}
	}
	if RabbitConn != nil {
		err := RabbitConn.Close()
		if err != nil {
			return err
		}
	}

	log.Println("RabbitMQ connection closed")
	return nil
}

func SendMessageToQueue(queueName string, messageBody string) error {
	if RabbitChannel == nil {
		return fmt.Errorf("RabbitMQ channel is nil")
	}

	_, err := RabbitChannel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Error.Printf("failed to declare queue: %v", err)
		return fmt.Errorf("failed to declare queue: %v", err)
	}

	err = RabbitChannel.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte(messageBody),
			DeliveryMode: amqp.Persistent,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}

	return nil
}
