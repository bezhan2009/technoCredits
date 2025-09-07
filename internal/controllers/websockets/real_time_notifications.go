package websockets

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"technoCredits/internal/controllers/middlewares"
	"technoCredits/pkg/brokers"
	"technoCredits/pkg/errs"
	"technoCredits/pkg/logger"
	"time"
)

func RealTimeNotificationReader(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error.Printf("WebSocket upgrade error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": errs.ErrSomethingWentWrong.Error()})
		return
	}
	defer conn.Close()

	queueName := fmt.Sprintf("user_%d_queue", userID)
	err = createDurableQueue(queueName)
	if err != nil {
		logger.Error.Printf("Failed to create queue for user %d: %v", userID, err)
		return
	}

	welcomeMsg := WebSocketMessage{
		Type:    "info",
		Payload: "Ждем сообщении об ваших долах :)",
	}
	conn.WriteJSON(welcomeMsg)

	//consumerName := fmt.Sprintf("user_%d_ws_consumer", userID)

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
		logger.Error.Printf("Failed to consume from queue %s: %v", queueName, err)
		return
	}

	for {
		select {
		case msg, ok := <-msgs:
			if !ok {
				logger.Info.Printf("Message channel closed for user %d", userID)
				return
			}

			wsMsg := WebSocketMessage{
				Type:    "message",
				Payload: string(msg.Body),
			}

			err = conn.WriteJSON(wsMsg)
			if err != nil {
				logger.Error.Printf("Failed to send message to user %d: %v", userID, err)

				msg.Nack(false, true)

				return
			}

			msg.Ack(false)
		case <-time.After(30 * time.Second):
			pingMsg := WebSocketMessage{
				Type:    "ping",
				Payload: "connection alive",
			}

			err = conn.WriteJSON(pingMsg)
			if err != nil {
				logger.Error.Printf("WebSocket ping failed for user %d: %v", userID, err)

				return
			}
		}
	}
}
