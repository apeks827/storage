package main

import (
	"fmt"

	"github.com/apeks827/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("it works", st)
}
