package config

import (
	"context"
	"github.com/jackc/pgx/v5"
	"os"
)

func ConfigureDatabase() *pgx.Conn {
	conn, err := pgx.Connect(context.TODO(), os.Getenv("DATABASE_DSN"))
	if err != nil {
		panic(err)
	}
	return conn
}
