package database

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectToDB() (*gorm.DB, error) {
	host := "host=localhost user=postgres password=postgres dbname=dojonov port=5432 sslmode=disable"

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		DSN: host,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Could not connect to database: %+v", err)
		return nil, err
	}

	log.Info("Database connected")

	return gormDB, err
}
