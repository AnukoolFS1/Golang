package main

import (
	"fmt"
	"context"
	"github.com/jackc/pgx/v5"
)

var ConnectionString = "postgres://postgres:bookingstore@localhost:5432/newdatabase"

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, ConnectionString)
	if err != nil {
		panic(err)
	}

	defer conn.Close(ctx)

	fmt.Println(conn)
	fmt.Println("Postgres is connected")
}