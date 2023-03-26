package postgres_client

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	instance *sql.DB
	once     sync.Once
)

const (
	DB_USER     = "priya"
	DB_PASSWORD = "d2xMromKh1Ay"
	DB_NAME     = "neondb"
	DB_HOST     = "ep-plain-base-359535.us-east-2.aws.neon.tech"
)

func SetupPostgres() *sql.DB {

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=%s", DB_USER, DB_PASSWORD, DB_NAME, DB_HOST, "require")
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if dbConn.Ping() != nil {
		log.Fatal(err)
	}

	return dbConn
}

func GetPostgresClient() *sql.DB {
	once.Do(func() {
		instance = SetupPostgres()
	})

	return instance
}
