package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

var db *gorm.DB
var err error

func Database() *gorm.DB {
	if db == nil {
		db, err = gorm.Open(postgres.Open(os.Getenv("DATABASE")), &gorm.Config{})

		if err != nil {
			log.Panic("failed to connect database")
		}

		log.Info("database connected")

		return db
	}

	return db
}
