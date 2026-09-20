package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

var ConnectionString = "postgres://postgres:bookingstore@localhost:5432/bookstore"

func main() {
	ctx := context.Background()

	db, err := pgxpool.New(ctx, ConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping(ctx) // Sends a lightweight query to the database and waits for a response, just to confirm the connection actually works.
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to PostgreSQL")

	_, err = db.Exec(
		ctx,
		`INSERT INTO books (title, author, price, stock)
		 VALUES ($2, $1, $3, $4)`,
		"The Go Programming Language",
		"Alan Donovan",
		799.00,
		10,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Book inserted")
}
// "NAMES IDENTIFIER"
// 'String Values'