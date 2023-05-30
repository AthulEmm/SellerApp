package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

// Database configuration
const (
	dbUsername = "root"
	dbPassword = "1234"
	dbHost     = "localhost"
	dbPort     = "3306"
	dbName     = "order_management"
)

// Global database connection
var (
	db   *sql.DB
	once *sync.Once = &sync.Once{}
)

func GetConnection() (*sql.DB, error) {
	var (
		err error
	)
	once.Do(func() {
		// Create a MySQL database connection
		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName))
		if err != nil {
			log.Fatal(err)
			return
		}
		// defer db.Close()
		// Test the database connection
		err = db.Ping()
		if err != nil {
			log.Fatal(err)
			return
		}
	})
	return db, err
}
