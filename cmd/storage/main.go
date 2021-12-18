package main

import (
	"fmt"
	"log"

	"github.com/apeks827/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	file, err := st.Upload("test.txt", []byte{"hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File upploaded", file)
}
