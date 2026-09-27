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

	// _, err = db.Exec(
	// 	ctx,
	// 	`INSERT INTO books (title, author, price, stock)
	// 	 VALUES ($2, $1, $3, $4)`,
	// 	"The Go Programming Language",
	// 	"Alan Donovan",
	// 	799.00,
	// 	10,
	// )

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Book inserted")
	// var bookID int

	// err = db.QueryRow(
	// 	ctx,
	// 	`INSERT INTO books (title, author, price, stock)
	//  VALUES ($1, $2, $3, $4)
	//  RETURNING id`,
	// 	"Harry Potter",
	// 	"JK Rowling",
	// 	89.00,
	// 	50,
	// ).Scan(&bookID)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Inserted book with ID:", bookID)

	var (
		id     int
		title  string
		author string
		price  float32
		stock  int
	)

	err = db.QueryRow(
		ctx,
		`SELECT id, title, author, price, stock FROM books WHERE id = $1`,
		6,
	).Scan(&id, &title, &author, &price, &stock)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(id)
	fmt.Println(title)
	fmt.Println(author)
	fmt.Println(price)
	fmt.Println(stock)

	row, err := db.Query(ctx, "SELECT id, title, author, price, stock FROM books;")

	fmt.Println("q", row.FieldDescriptions())

	for row.Next() {

		fmt.Println(row.Values())

		var (
			id     int
			title  string
			author string
			price  float32
			stock  int
		)

		err := row.Scan(&id, &title, &author, &price, &stock)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("id: %v, title: %v, author: %v, price: %v, stock: %v \n", id, title, author, price, stock)
	}

	if err := row.Err(); err != nil {
		log.Fatal(err)
	}

	defer row.Close()

}

// "NAMES IDENTIFIER"
// 'String Values'
