package rabbitmq

import (
	utils "github.com/mdas-ds2/mdas-api-g3/src/generic/infrastructure/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitQueue struct {
	channel *amqp.Channel
	queue   amqp.Queue
}

func (q RabbitQueue) Publish(payload []byte) {
	msg := amqp.Publishing{
		ContentType:  "application/json",
		DeliveryMode: amqp.Persistent,
		Priority:     0,
		Body:         payload,
	}

	err := q.channel.Publish(
		"",           // exchange
		q.queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		msg,
	)
	utils.FailOnError(err, "Failed to publish a message")
}

func (q RabbitQueue) Name() string {
	return q.queue.Name
}
