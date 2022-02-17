package server

import (
	"net/http"

	"github.com/blattaria7/go-template/internal/storage"
)

type Server struct {
	client  *http.Client
	storage storage.Storager
}

func NewServer(client *http.Client, storage storage.Storager) *Server {
	return &Server{
		client:  client,
		storage: storage,
	}
}

// .. here you initialize your server interface

// .. here you initialize your server methods
