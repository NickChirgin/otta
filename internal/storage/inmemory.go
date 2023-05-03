package storage

import (
	"fmt"
	"log"

	"github.com/hashicorp/go-memdb"
	"github.com/nickchirgin/otta/pkg/hasher"
)

type MemoryDB struct {
	DB *memdb.MemDB
}

type Row struct {
	id int
	shortURL string
	url string
}

var (
	MemDB *MemoryDB
	id int
)
func init() {
	id = 1
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"urls": {
				Name: "urls",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "Id"},
					},
					"tinyURL": {
						Name:    "tinyURL",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "TinyURL"},
					},
					"fullURL": {
						Name:    "fullURL",
						Unique:  false,
						Indexer: &memdb.IntFieldIndex{Field: "FullURL"},
					},
				},
			},
		},
	}
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		log.Fatalf("Error %v while connecting to in-memory database", err)
	}
	MemDB = &MemoryDB{DB: db}
	fmt.Println("In-memory db ready to use")
}

func (m *MemoryDB) ShortUrl(url string) string {
	tinyURL, err := m.URLExist(url)
	if err != nil {
		tinyURL = hasher.HashURL(id)
		err = m.AddURL(url, tinyURL)
		if err != nil {
			log.Fatalf("Error while adding to inmemory db %v", err)
		}
	}
	return tinyURL
}

func (m *MemoryDB) FullURL(hashedURL string) string {
	txn := MemDB.DB.Txn(false)
	defer txn.Abort()
	raw,err := txn.First("urls", "tinyURL", hashedURL)
	if err != nil {
		log.Fatalf("No short URL for this site %v", err)
	}
	return raw.(Row).url
}

func (m *MemoryDB) AddURL(url, shortURL string) error {
	row := Row{id: id, url: url, shortURL: shortURL}
	txn := MemDB.DB.Txn(true)
	err := txn.Insert("urls", row)
	if err != nil {
		return err
	}
	txn.Commit()
	id++
	return nil
}

func (m *MemoryDB) URLExist(url string) (string, error) {
	txn := MemDB.DB.Txn(false)
	defer txn.Abort()
	raw, err := txn.First("urls", "fullURL", url)
	if err != nil {
		return "", err
	}
	return raw.(Row).shortURL, nil
}