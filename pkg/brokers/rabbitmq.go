package brokers

import (
	"github.com/streadway/amqp"
	"log"
	"technoCredits/internal/app/models"
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
