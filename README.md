# mdas-api-ezequiel_gomez
MDAS - Diseño de software II - DDD - Reto individual

### Contributors:

- ezequiel.gomez@students.salle.url.edu

## Installation steps

_This project requires Go +1.13 and Go module support._

1. Clone the repository
   git clone https://github.com/mdas-ds2/mdas-api-ezequiel_gomez.git .

2. Run the application

```
docker-compose up --build
```

3. Test execution

```
go test ./...
```

4. After application started a webserver will be held on port 5001, so that you can:

- get types via http request: curl http://localhost:5001/pokemon-types\?name\=charizard
- add pokemon favorite(squirtle) to user(1234): curl -X POST -H "UserId:1234" -d '{"pokemonId": "squirtle"}' http://localhost:5001/favorite-pokemon/
- get pokemon details : http://localhost:5001/pokemon/?id=25
