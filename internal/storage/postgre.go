package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/nickchirgin/otta/pkg/hasher"
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
	tinyURL, err := p.URLExist(url)
	if err != nil {
		id, err := p.LastID()
		if err != nil {
			log.Fatalf("Error while finding last id in postgre %v", err)
		}
		tinyURL = hasher.HashURL(int(id))
		err = p.AddURL(url, tinyURL)
		if err != nil {
			log.Fatalf("Error while adding data to postgre %v", err)
		}
	}
	return tinyURL
}

func (p *PostgreSQL) FullURL(shortUrl string) string {
	stmt, err := p.DB.Prepare("SELECT url FROM urls WHERE shorturl=$1")
	if err != nil {
		log.Fatalf("Error while querying postgre %v", err)
	}
	var url string
	defer stmt.Close()
	err = stmt.QueryRow(shortUrl).Scan(url)
	if err != nil {
		log.Fatalf("Error while executing query %v", err)
	}
	return url
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

func (p *PostgreSQL) AddURL(url, shorturl string) error {
	stmt, err := p.DB.Prepare("INSERT INTO urls (url, shorturl) VALUES ($1, $2);") 
	if err != nil {
		return err
	}
	defer stmt.Close()
	stmt.QueryRow(url, shorturl)
	return nil
}