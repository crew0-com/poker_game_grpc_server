package main

import (
	"database/sql"
	"log"
	"poker/internal/config"

	_ "github.com/lib/pq"
	db "poker/database/sqlc"
)

func main() {
	c, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	conn, err := sql.Open(c.DBDriver, c.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	store := db.New(conn)
	log.Printf("store: %v", store)
}
