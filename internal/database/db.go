package database

import (
	"fmt"
	"log"

	config "task-management-service/config"
	model "task-management-service/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(cfg *config.Config) *gorm.DB {
	url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=UTC",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName, cfg.Database.Port, cfg.Database.SSLMode)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	return db
}

func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func AutoMigrate(database *gorm.DB) error {
	return database.AutoMigrate(&model.Task{}, &model.User{})
}
