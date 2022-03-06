package rabbitmq

import (
	utils "github.com/mdas-ds2/mdas-api-g3/src/generic/infrastructure/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func (rabbit *RabbitMQ) NewQueue(queueName string) RabbitQueue {
	queue, err := rabbit.Channel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	utils.FailOnError(err, "Failed to declare a queue")

	return RabbitQueue{rabbit.Channel, queue}
}

func New() *RabbitMQ {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")

	return &RabbitMQ{conn, ch}
}
