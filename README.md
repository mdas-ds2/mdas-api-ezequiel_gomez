# mdas-api-ezequiel_gomez
MDAS - Diseño de software II - DDD - Reto individual

### Contributors:

- ezequiel.gomez@students.salle.url.edu

## Installation steps

_This project requires Go +1.13 and Go module support._

1. Clone the repository
   git clone https://github.com/mdas-ds2/mdas-api-ezequiel_gomez.git .

2. Run the rabbitMQ container

```
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
```

3. Run the main appplication once the above container is running:

```
go run . -getPokemonTypes pikachu
```

4. Test execution

```
go test ./...
```

5. After application started a webserver will be held on port 5001, so that you can:

- get types via http request: curl http://localhost:5001/pokemon-types\?name\=charizard
- add pokemon favorite(squirtle) to user(1234): curl -X POST -H "UserId:1234" -d '{"pokemonId": "25"}' http://localhost:5001/favorite-pokemon/
- get pokemon details : http://localhost:5001/pokemon/?id=25
