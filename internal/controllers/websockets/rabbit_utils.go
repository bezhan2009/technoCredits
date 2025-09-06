package websockets

import (
	"fmt"
	"github.com/gorilla/websocket"
	"technoCredits/pkg/brokers"
	"technoCredits/pkg/logger"
	"time"
)

func createDurableQueue(queueName string) error {
	_, err := brokers.RabbitChannel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Error.Printf("failed to declare durable queue %s: %v", queueName, err)
		return fmt.Errorf("failed to declare durable queue %s: %v", queueName, err)
	}

	return nil
}

func readAllMessagesFromQueue(queueName string, conn *websocket.Conn) error {
	msgs, err := brokers.RabbitChannel.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Error.Printf("failed to consume from queue %s: %v", queueName, err)
		return fmt.Errorf("failed to consume from queue %s: %v", queueName, err)
	}

	timeout := time.After(5 * time.Second)
	for {
		select {
		case msg, ok := <-msgs:
			if !ok {
				return nil
			}

			wsMsg := WebSocketMessage{
				Type:    "message",
				Payload: string(msg.Body),
			}

			err = conn.WriteJSON(wsMsg)
			if err != nil {
				msg.Nack(false, true)
				return fmt.Errorf("failed to send message to client: %v", err)
			}

			msg.Ack(false)
			logger.Info.Printf("Message sent and acknowledged: %s", string(msg.Body))

		case <-timeout:
			logger.Info.Printf("Finished reading messages from queue: %s", queueName)
			return nil
		}
	}
}
