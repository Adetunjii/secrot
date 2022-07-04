package main

import (
	gRPC "github.com/Adetunjii/secrot/internal/adapters/framework/in/grpc"
	"github.com/Adetunjii/secrot/internal/application/api"
	"github.com/Adetunjii/secrot/internal/application/domain"
	"github.com/Adetunjii/secrot/utils"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("couldn't load config: %v", err)
	}

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("error %v", err)
	}

	core := domain.New()

	application := api.NewApplication(config.DB, core)

	grpcAdapter := gRPC.NewAdapter(application)
	grpcAdapter.Run()
}
