package app

import (
	"log"

	config "task-management-service/configs"
)

func Run() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

}
