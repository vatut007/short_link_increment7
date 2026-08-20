package config

import (
	"flag"
	"os"
)

func ParseFlags() {
	flag.StringVar(&ConfigApp.ServerAddress, "a", "localhost:8080", "HTTP server address (e.g. localhost:8888)")
	flag.StringVar(&ConfigApp.BaseURL, "b", "http://localhost:8080", "Base URL for short links (e.g. http://localhost:8000)")
	flag.Parse()
	if envServerAddress := os.Getenv("SERVER_ADDRESS"); envServerAddress != "" {
		ConfigApp.ServerAddress = envServerAddress
	}
	if envBaseURL := os.Getenv("BASE_URL"); envBaseURL != "" {
		ConfigApp.BaseURL = envBaseURL
	}
}
