package main

import (
	"log"
	"returnbook/internal/storage"
	"returnbook/internal/workflow36"
)

func main() {
	s, e := storage.Open("returnbook.db")
	if e != nil {
		log.Fatal(e)
	}
	defer s.Close()
	_ = workflow36.New(s)
}
