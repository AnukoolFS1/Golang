package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

var ConnectionString = "postgres://postgres:bookingstore@localhost:5432/newdatabase"

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, ConnectionString)
	if err != nil {
		panic(err)
	}

	_, err = conn.Exec(ctx, `DROP TABLE newtable;`)

	if err != nil {
		panic(err)
	}

	// fmt.Println(string(value))

	defer conn.Close(ctx)

	fmt.Println("Postgres is connected")
	fmt.Println("Table is created")
}
// "NAMES IDENTIFIER"
// 'String Values'