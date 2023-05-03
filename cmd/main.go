package main

import (
	"fmt"

	"github.com/nickchirgin/otta/internal/storage"
	"github.com/nickchirgin/otta/pkg/hasher"
)

type Server struct {
	Data storage.IStorage
	Age int
}

func main() {
	fmt.Println(hasher.HashURL(125))
}