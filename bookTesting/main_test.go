package bookTesting

import (
	"book/config"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	DBmanager *DBManager
)

func TestMain(m *testing.M) {
	cfg := config.Load("")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
	cfg.Postgres.Host,
	cfg.Postgres.Port,
	cfg.Postgres.User,
	cfg.Postgres.Password,
	cfg.Postgres.Database,
	)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect postgres: %v", err)
	}
	DBmanager = NewDbmanager(db)
	os.Exit(m.Run())
}
