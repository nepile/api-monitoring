package main

import (
	"log"

	"github.com/nepile/api-monitoring/config"
	"github.com/nepile/api-monitoring/database"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	database.Connect(cfg.DatabaseURL)
}
