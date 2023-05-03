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

func (p *PostgreSQL) ShortUrl(url string) string {

}

func (p *PostgreSQL) FullURL(shortUrl string) string {
	return ""
}

func (p *PostgreSQL) LastID() (int64, error) {
	var id int64
	result, err := p.DB.Exec("SELECT currval(pg_get_serial_sequence('urls', 'id'))")
	if err != nil {
		return 0, err
	}
	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	} 
	return id, nil
}

func (p *PostgreSQL) URLExist(url string) (string, error) {
	stmt, err := p.DB.Prepare("SELECT shorturl FROM urls WHERE url=$1;")
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	var shortURL string
	err = stmt.QueryRow(url).Scan(shortURL)
	if err != nil {
		return "", err
	}
	return shortURL, nil	
}