package main

import (
	"log"

	"github.com/pets-shelters/backend-svc/configs"
	"github.com/pets-shelters/backend-svc/internal/app"
)

func main() {
	cfg, err := configs.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
