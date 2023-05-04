package main

import (
	"context"
	"log"
	"net"

	"github.com/nickchirgin/otta/internal/helpers"
	"github.com/nickchirgin/otta/internal/storage"
	"github.com/nickchirgin/otta/proto"
	"google.golang.org/grpc"
)
type server struct {
	proto.UnimplementedUrlServiceServer
	DB storage.IStorage
}

func main() {
	db := helpers.ChooseDB()
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterUrlServiceServer(s, &server{DB: db})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) TinyURL(ctx context.Context, req *proto.URL) (*proto.HashedURL, error) {
	url := req.GetFullURL()
	shortURL := s.DB.ShortUrl(url)
	return &proto.HashedURL{ShortURL: shortURL}, nil
}

func (s *server) FullURL(ctx context.Context, req *proto.HashedURL) (*proto.URL, error) {
	shortURL := req.GetShortURL()
	URL, err := s.DB.FullURL(shortURL)
	if err != nil {
		return nil, err
	}
	return &proto.URL{FullURL: URL}, nil
}

