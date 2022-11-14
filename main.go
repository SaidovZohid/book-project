package main

import (
	"fmt"
	"log"

	bo "book/bookTesting"
	"book/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	DBmanager *bo.DBManager
)

func main() {
	cfg := config.Load(".")
	fmt.Println(cfg)

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Database)
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