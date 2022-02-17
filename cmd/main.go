package main

import (
	"github.com/blattaria7/go-template/internal/server"
	"github.com/blattaria7/go-template/internal/storage"
	"net/http"
)

func main() {
	// initialize your storage and storage's variables
	items := make(map[string]string)
	s := storage.NewStorage(items)

	// initialize your server and server's variables
	client := http.Client{} // mock
	_ = server.NewServer(&client, s)

	// ... here you start your server
}
