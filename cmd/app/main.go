package main

import (
	"log"

	"github.com/pets-shelters/backend-svc/configs"
)

func main() {
	// Configuration
	cfg, err := configs.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
