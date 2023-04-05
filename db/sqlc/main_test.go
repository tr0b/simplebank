package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"

	"github.com/tr0b/simplebank/internal/projectpath"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	envErr := godotenv.Load(projectpath.Root + "/.env")
	path, _ := os.Getwd()
	log.Println(path)
	if envErr != nil {
		log.Fatal("Error loading environment variables file:", envErr)
	}
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	port := os.Getenv("POSTGRES_PORT")
	sslMode := os.Getenv("SSL_MODE")
	database := os.Getenv("POSTGRES_DATABASE")

	dbDriver := "postgres"
	dbSource := "postgresql://" + user + ":" + password + "@localhost:" + port + "/" + database + "?sslmode=" + sslMode
	var err error

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())

}
