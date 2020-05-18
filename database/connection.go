package database

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

type params struct {
	dbms     string
	username string
	password string
	dbName   string
	dbHost   string
}

func init() {

	var config params

	valuesEnv := godotenv.Load()
	if valuesEnv != nil {
		log.Fatal(valuesEnv)
	}

	config.dbms = os.Getenv("DBMS")
	config.username = os.Getenv("db_user")
	config.password = os.Getenv("db_pass")
	config.dbName = os.Getenv("db_name")
	config.dbHost = os.Getenv("db_host")
	conn, err := gorm.Open(config.dbms, config.username+":"+config.password+"@/"+config.dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db = conn
}

func DBConnect() *gorm.DB {
	return db
}

/* Go closes the DB connection for each client
func DBClose() {
	if db != nil {
		db.Close()
	}
}*/
