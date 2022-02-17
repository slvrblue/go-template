package main

import "github.com/blattaria7/go-template/internal/storage"

func main() {
	items := make(map[string]string)
	_ = storage.NewStorage(items)
}
