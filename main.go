package main

import (
	rabbitMQ "github.com/mdas-ds2/mdas-api-g3/src/generic/infrastructure/msg-broker/rabbitmq"
	webServer "github.com/mdas-ds2/mdas-api-g3/src/generic/infrastructure/web-server"
	pokemonTypesCommands "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/infrastructure/commands"
	pokemonTypesControllers "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/infrastructure/controllers"
	pokemonControllers "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/infrastructure/controllers"
	pokemonSubscribers "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/infrastructure/subscribers"
	userControllers "github.com/mdas-ds2/mdas-api-g3/src/users/user/infrastructure/controllers"
)

func main() {
	rabbit := rabbitMQ.New()
	defer rabbit.Connection.Close()
	defer rabbit.Channel.Close()

	pokemonQueue := rabbit.NewQueue("pokemon-marked-as-favorite")
	pokemonQueueConsumer := rabbitMQ.CreateRabbitMQConsumer(*rabbit.Channel, pokemonQueue)

	pokemonSubscribers.CreateUpdatePokemonOnMarkedAsFavoriteSubscriber(pokemonQueueConsumer).Run()

	pokemonTypesCommands.NewGetTypesByPokemonName().Run()

	server := webServer.Create()
	server.Register(pokemonTypesControllers.CreateGetTypesByPokemonName())
	server.Register(userControllers.CreateAddFavoritePokemonController(pokemonQueue))
	server.Register(pokemonControllers.CreateGetPokemonDetailsController())
	server.Listen(5001)
}
