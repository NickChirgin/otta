package storage

import (
	"fmt"
	"log"

	"github.com/hashicorp/go-memdb"
)

type MemoryDB struct {
	DB *memdb.MemDB
}

var MemDB *MemoryDB
func init() {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"url": {
				Name: "url",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "HashedURL"},
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

func (m *MemoryDB) GetShortUrl(url string) string {
	return ""
}

func (m *MemoryDB) GetFullURL(hashedURL string) string {
	return ""
}