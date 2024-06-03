package postgres

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func ConnectToPostgres(postgresURL string) *sql.DB {
	db, err := sql.Open("postgres", postgresURL)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
