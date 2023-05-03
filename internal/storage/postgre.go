package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/stdlib"
)

type PostgreSQL struct {
	DB *sql.DB
}
var (
	Postgre *PostgreSQL
	username = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	schema   = os.Getenv("DB_NAME")
	host = os.Getenv("DB_HOST")
)

func init() {
	connInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host,	username, password, schema)
	var err error
	conn, err := sql.Open("pgx", connInfo)
	if err != nil {
		log.Fatalf("Error %v while connecting to postgre db", err)
	}
	if err = conn.Ping(); err != nil {
		log.Fatalf("Error %v while pinging postgre db", err)
	} 
	log.Println("Postgre database ready to accept connections")
	Postgre = &PostgreSQL{DB: conn}
}

func (p *PostgreSQL) GetShortUrl(url string) string {
	return ""
}

func (p *PostgreSQL) GetFullURL(shortUrl string) string {
	return ""
}