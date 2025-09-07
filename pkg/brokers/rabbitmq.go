package brokers

import (
	"fmt"
	"sync"
	"technoCredits/internal/security"
	"time"

	"github.com/streadway/amqp"
	"technoCredits/internal/app/models"
	"technoCredits/pkg/logger"
)

var (
	RabbitConn     *amqp.Connection
	RabbitChannel  *amqp.Channel
	rabbitMutex    sync.Mutex
	reconnectDelay = 2 * time.Second
)

func ConnectToRabbitMq(params models.RabbitParams) error {
	rabbitMutex.Lock()
	defer rabbitMutex.Unlock()

	return connect(params)
}

func connect(params models.RabbitParams) error {
	var err error

	if RabbitConn != nil {
		RabbitConn.Close()
	}

	RabbitConn, err = amqp.Dial(params.URLConn)
	if err != nil {
		return fmt.Errorf("failed to dial RabbitMQ: %v", err)
	}

	RabbitChannel, err = RabbitConn.Channel()
	if err != nil {
		RabbitConn.Close()
		return fmt.Errorf("failed to open channel: %v", err)
	}

	go func() {
		<-RabbitConn.NotifyClose(make(chan *amqp.Error))
		logger.Info.Println("RabbitMQ connection closed, attempting to reconnect...")
		reconnect(params)
	}()

	logger.Info.Println("Connected to RabbitMQ successfully")
	return nil
}

func reconnect(params models.RabbitParams) {
	time.Sleep(reconnectDelay)

	rabbitMutex.Lock()
	defer rabbitMutex.Unlock()

	for {
		err := connect(params)
		if err == nil {
			logger.Info.Println("Reconnected to RabbitMQ successfully")
			return
		}

		logger.Error.Printf("Failed to reconnect: %v. Retrying in %v", err, reconnectDelay)
		time.Sleep(reconnectDelay)
	}
}

func IsConnected() bool {
	rabbitMutex.Lock()
	defer rabbitMutex.Unlock()

	return RabbitConn != nil && !RabbitConn.IsClosed() &&
		RabbitChannel != nil
}

func EnsureConnection(params models.RabbitParams) error {
	rabbitMutex.Lock()
	defer rabbitMutex.Unlock()

	if RabbitConn != nil && !RabbitConn.IsClosed() &&
		RabbitChannel != nil {
		return nil
	}

	return connect(params)
}

func SendMessageToQueue(queueName string, messageBody string) error {
	if err := EnsureConnection(security.AppSettings.RabbitParams); err != nil {
		return fmt.Errorf("failed to ensure connection: %v", err)
	}

	rabbitMutex.Lock()
	defer rabbitMutex.Unlock()

	_, err := RabbitChannel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
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

func CloseRabbitMQ() error {
	rabbitMutex.Lock()
	defer rabbitMutex.Unlock()

	if RabbitChannel != nil {
		RabbitChannel.Close()
	}
	if RabbitConn != nil {
		return RabbitConn.Close()
	}
	return nil
}
