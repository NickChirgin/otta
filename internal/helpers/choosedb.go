package helpers

import (
	"os"

	"github.com/nickchirgin/otta/internal/storage"
)

func ChooseDB() storage.IStorage {
	if len(os.Args) > 1 {
		memory := os.Args[1]
		if memory == "-memory"	{
			return storage.MemDB
		}
	}
	return storage.Postgre
}