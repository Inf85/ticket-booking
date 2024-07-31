package db

import (
	"fmt"
	"github.com/Inf85/ticket-booking/config"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)
import _ "gorm.io/gorm"

func Init(config *config.EnvConfig, DBMigrator func(db *gorm.DB) error) *gorm.DB {
	uri := fmt.Sprintf(`host=%s user=%s dbname=%s password=%s sslmode=%s port=5432`,
		config.DBHost, config.DBUser, config.DBName, config.DBPassword, config.DBSSLMode)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Unable to connect to Database %e", err)
	}

	log.Info("Connected to Database")

	if err := DBMigrator(db); err != nil {
		log.Fatalf("Unable to migrate tables to Database %e", err)
	}
	return db
}
