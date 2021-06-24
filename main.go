package main

import (
	"fmt"
	"github.com/forderation/hospital-information-system/db"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load(fmt.Sprintf(".env.%s", "development"))
	if err != nil {
		panic(err.Error())
	}
	/**
	Try Connect Mysql DB
	*/
	configMysql := db.Config{
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		DBName:   os.Getenv("MYSQL_DBNAME"),
	}
	database := db.ConnectMysql(configMysql)
	db.Migrate(database)
	db.SeedUser(database, 5)
}
