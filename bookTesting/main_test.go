package bookTesting

import (
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

var (
	DBmanager *DBManager
	host      = "localhost"
	port      = 5432
	user      = "postgres"
	password  = "1234"
	dbname    = "book"
)

func TestMain(m *testing.M) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect postgres: %v", err)
	}
	DBmanager = NewDbmanager(db)
	os.Exit(m.Run())
}
