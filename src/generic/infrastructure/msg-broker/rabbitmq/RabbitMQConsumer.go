package rabbitmq

import (
	utils "github.com/mdas-ds2/mdas-api-g3/src/generic/infrastructure/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

type AmqpBus = <-chan amqp.Delivery

type RabbitMQConsumer struct{}

func CreateRabbitMQConsumer(channel amqp.Channel, queue RabbitQueue) AmqpBus {
	msgs, err := channel.Consume(
		queue.Name(), // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	utils.FailOnError(err, "Failed to register a consumer")

	return msgs
}
