package main

import (
	"fmt"
	"github.com/forderation/hospital-information-system/db"
	"github.com/forderation/hospital-information-system/service"
	"github.com/joho/godotenv"
	"log"
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
	//db.Migrate(database)
	//db.SeedUser(database, 5)
	router := service.InitRoute(database.DB)
	err = router.Run("127.0.0.1:80")
	if err != nil {
		log.Fatal(err.Error())
	}
}
