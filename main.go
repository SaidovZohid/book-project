package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	bo "book/bookTesting"
)

var (
	DBmanager *bo.DBManager
	host      = "localhost"
	port      = 5432
	user      = "postgres"
	password  = "1234"
	dbname    = "book"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect postgres: %v", err)
	}
	DBmanager = bo.NewDbmanager(db)
	server := NewServer(DBmanager)

	err = server.Run(":8080")
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}