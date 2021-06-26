package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type DB struct {
	DB *gorm.DB
}

func ConnectMysql(config Config) *DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.DBName,
	)
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}
	if os.Getenv("USING_LOG") == "true" {
		loggerDb := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
				Colorful:                  false,
			},
		)
		gormConfig.Logger = loggerDb
	}
	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		panic(err.Error())
	}
	newDb := DB{
		DB: db,
	}
	return &newDb
}
