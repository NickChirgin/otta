package main

import (
	"github.com/nickchirgin/otta/internal/storage"
)

type Server struct {
	Data storage.IStorage
	Age int
}

func main() {
	server1 := Server{Data: storage.Postgre, Age: 15}
	server1.Data.GetFullURL("")
	server2 := Server{Data: storage.MemDB, Age: 10}
	server2.Data.GetFullURL("")
}