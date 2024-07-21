package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectDB(conn string) (*pgx.Conn, error) {
	connected, err := pgx.Connect(context.Background(), conn)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error occured when connecting to db %v\n", err)
		return nil, err
	}

	return connected, nil
}
