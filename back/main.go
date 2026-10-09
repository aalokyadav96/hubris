package main

import (
	"log"

	"naevis/config"
	"naevis/infra"
	"naevis/server"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	cfg := config.InitConfig()
	app, err := infra.New(cfg)
	if err != nil {
		log.Fatalf("failed to initialize infra: %v", err)
	}

	if err := server.Run(app); err != nil {
		log.Fatalf("server exited with error: %v", err)
	}
}
