# Reverse Youtube Search in GO

## Locally

1. Copy `.env.example` => `.env`, alternatively run `make init`.
2. Populate `YOUTUBE_API_KEY` env variable within `.env` file.
3. Fetch dependencies `go get`
4. Start the server `go run main.go`, alternatively run `make run`.

## Docker

1. Follow the steps to populate `YOUTUBE_API_KEY` from above.
2. Run the server `docker compose up -d`
3. Make sure port mapping within `docker-compose.yml` is corresponding to the value set in `.env`
4. To forcefully recreate container run `docker compose up -d --build --force-recreate`
5. Alternatively run `make run_docker`

## Test

Server exposes an endpoint: `***/v1/youtube/search` accepting a search param, fe: `http://127.0.0.1:8080/v1/youtube/search?q=ibet365`
