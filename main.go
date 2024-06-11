package main

import (
	"fmt"
	"go-reverse-youtube-search/env"
	"go-reverse-youtube-search/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	env.Init()

	router := http.NewServeMux()
	router.HandleFunc("/v1/youtube/search", handlers.YoutubeSearchHandler)

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv(env.PORT)),
		Handler: router,
	}

	log.Printf("Listening on port: %s", os.Getenv(env.PORT))

	if err := server.ListenAndServe(); err != nil {
		panic(err.Error())
	}
}
