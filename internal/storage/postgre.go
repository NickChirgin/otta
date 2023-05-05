package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/nickchirgin/otta/pkg/hasher"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PostgreSQL struct {
	DB *sql.DB
	id int
}
var (
	Postgre *PostgreSQL
	username = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	schema   = os.Getenv("DB_NAME")
	host = "10.5.0.3" 
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
	Postgre = &PostgreSQL{DB: conn, id: 1}
}

func (p *PostgreSQL) ShortUrl(url string) string {
	tinyURL, err := p.URLExist(url)
	if err != nil {
		tinyURL = hasher.HashURL(p.id)
		err = p.AddURL(url, tinyURL)
		if err != nil {
			log.Fatalf("Error while adding data to postgre %v", err)
		}
		p.id++
	}
	return tinyURL
}

func (p *PostgreSQL) FullURL(shortUrl string) (string, error) {
	stmt, err := p.DB.Prepare("SELECT url FROM urls WHERE shorturl=$1")
	if err != nil {
		return "", err
	}
	var url string
	defer stmt.Close()
	err = stmt.QueryRow(shortUrl).Scan(&url)
	if err != nil {
		return "", status.Error(codes.NotFound, "Row doesnt exist")
	}
	return url, nil
}

func (p *PostgreSQL) LastID() (int64, error) {
	var id int64
	result, err := p.DB.Exec("SELECT currval('urls', 'id')")
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
	err = stmt.QueryRow(url).Scan(&shortURL)
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