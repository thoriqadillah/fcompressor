package db

import (
	"fcompressor/env"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var instance *gorm.DB

func DB() *gorm.DB {
	return instance
}

func Open() {
	var (
		host     = env.Get("DB_HOST").String("localhost")
		port     = env.Get("DB_PORT").Int(5432)
		username = env.Get("DB_USERNAME").String("postgres")
		password = env.Get("DB_PASSWORD").String()
		db       = env.Get("DB_NAME").String("postgres")
	)

	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  true,          // Disable color
		},
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, username, password, db, port)
	_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger,
	})

	if err != nil {
		log.Panic("DB ERROR", err)
	}

	instance = _db
}
