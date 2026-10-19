package main

import (
	"log"
	"net/http"

	"github.com/vatut007/short_link_increment7/cmd/shortener/config"
	"github.com/vatut007/short_link_increment7/handlers"
)

func main() {
	config.ParseFlags()
	main_router := handlers.NewRouter()
	if err := http.ListenAndServe(config.ConfigApp.ServerAddress, main_router); err != nil {
		log.Fatal(err)
	}
}
