package main

import (
	"log"

	"github.com/dessinyunyun/socialgo/internal/db"
	"github.com/dessinyunyun/socialgo/internal/repository"
)

const version = "1.1.0"

func main() {
	cfg := config{
		addr: ":8080",
		db: dbConfig{
			addr:         "postgres://postgres:adminpassword@localhost/social?sslmode=disable",
			maxOpenConns: 30,
			maxIdleConns: 30,
			maxIdleTime:  "15m",
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("database connection pool estabilished")

	repository := repository.NewRepository(db)

	app := &application{
		config:     cfg,
		repository: repository,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
