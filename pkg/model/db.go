package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func DB() {

	err := godotenv.Load("../config/.env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err.Error())
		return
	}
	config := fmt.Sprintf("%s:%s@%s(%s)/%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PW"),
		os.Getenv("MYSQL_NET"),
		os.Getenv("MYSQL_ADD"),
		os.Getenv("MYSQL_DB_NAME"),
	)

	db, err = sql.Open("mysql", config)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("Error could not ping database: %s\n", err.Error())
		return
	} else {
		fmt.Printf("DB pinged successfully\n")
	}

}
