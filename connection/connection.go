package connection

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var Db *sql.DB // DB Connection Manager

func ConnectDB() {
	errVar := godotenv.Load()
	// Check ENV Variables Load
	if errVar != nil {
		panic(errVar)
	}

	credentials := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SERVER"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	connection, err := sql.Open("mysql", credentials)
	// Check DB Connection
	if err != nil {
		panic(err)
	}

	Db = connection
}

func CloseDB() {
	Db.Close()
}
