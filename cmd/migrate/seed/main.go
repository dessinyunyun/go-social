package main

import (
	"log"

	"github.com/dessinyunyun/socialgo/internal/db"
	"github.com/dessinyunyun/socialgo/internal/repository"
)

func main() {
	addr := "postgres://postgres:adminpassword@localhost/social?sslmode=disable"

	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	repository := repository.NewRepository(conn)

	db.Seed(repository)
}
