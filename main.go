package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"

	"github.com/tr0b/simplebank/api"
	db "github.com/tr0b/simplebank/db/sqlc"
	"github.com/tr0b/simplebank/internal/projectpath"
)

const (
	serverAddress = "0.0.0.0:8080"
)

func main() {
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
	database := os.Getenv("POSTGRES_DB")

	dbDriver := "postgres"
	dbSource := "postgresql://" + user + ":" + password + "@localhost:" + port + "/" + database + "?sslmode=" + sslMode
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
