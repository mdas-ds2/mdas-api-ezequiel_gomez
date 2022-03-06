package main

import (
	"encoding/json"
	"log"

	webServer "github.com/mdas-ds2/mdas-api-g3/src/generic/infrastructure/web-server"
	pokemonTypesCommands "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/infrastructure/commands"
	pokemonTypesControllers "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/infrastructure/controllers"
	pokemonsControllers "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/infrastructure/controllers"
	usersControllers "github.com/mdas-ds2/mdas-api-g3/src/users/user/infrastructure/controllers"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	// Set connection and channel
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"pokemon", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Producer
	payload := map[string]int{"hello": 1}
	bodyJson, err := json.Marshal(payload)

	failOnError(err, "Failed on formatting payload to JSON format")

	msg := amqp.Publishing{
		ContentType:  "application/json",
		DeliveryMode: amqp.Persistent,
		Priority:     0,
		Body:         bodyJson,
	}
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		msg,
	)
	failOnError(err, "Failed to publish a message")

	// Consumer
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	// Main app
	pokemonTypesCommands.NewGetTypesByPokemonName().Run()

	server := webServer.Create()
	server.Register(pokemonTypesControllers.CreateGetTypesByPokemonName())
	server.Register(usersControllers.CreateAddFavoritePokemonController())
	server.Register(pokemonsControllers.CreateGetPokemonDetailsController())
	server.Listen(5001)
}
