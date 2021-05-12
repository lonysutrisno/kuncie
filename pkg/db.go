package pkg

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB
var err error

func init() {
	// Create an sql.DB and check for errors

	dbCon := os.Getenv("DB_CONNECTION")
	if dbCon == "" {
		dbCon = "root:123456@tcp(127.0.0.1:33061)/kuncie"
	}

	DB, err = sqlx.Open("mysql", dbCon)

	if err != nil {
		fmt.Println(err)
	}

	// Test the connection to the database
	err = DB.Ping()
	if err != nil {
		fmt.Println(err)
	}
	log.Println("connected to db")
}
