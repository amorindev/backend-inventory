package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)


var DB *sql.DB

func DBConnection() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")	
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if dbUser == "" || dbPass == "" || dbHost == ""|| dbName == ""{
		log.Fatal("One or more enviroments variables are not set")
	}

	sqlString := "host=%s user=%s password=%s dbname=%s port=%s sslmode=require"
	DSN := fmt.Sprintf(sqlString, dbHost, dbUser, dbPass, dbName, dbPort)

	var err error
	DB,err = sql.Open("postgres",DSN)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Connected to postgresql ...!")
}