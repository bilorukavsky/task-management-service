package app

import (
	config "task-management-service/config"
	db "task-management-service/internal/database"
	log "task-management-service/pkg/log"

	"github.com/joho/godotenv"
)

func Run() {
	err := godotenv.Load("../.env")
	logger := log.InitLogger()
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatalf("Error config load: %v", err)
	}

	database := db.Init(cfg)
	defer func() {
		if err := db.Close(database); err != nil {
			logger.Fatalf("Error closing the database: %v", err)
		}
	}()

	if err = db.AutoMigrate(database); err != nil {
		logger.Fatalf("Error automigrate of database: %v", err)
	}

	logger.Print("DB migrated successfull")

}
