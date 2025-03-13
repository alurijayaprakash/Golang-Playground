package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	Client *sqlx.DB
}

func NewDatabase() (*Database, error) {
	connectionStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_TABLE"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SSL_MODE"),
	)
	// here we will Connect to a database and verify with a ping.
	dbConn, err := sqlx.Connect("postgres", connectionStr)
	if err != nil {
		return &Database{}, fmt.Errorf("could not connect to the database: %w", err)
	}
	return &Database{
		Client: dbConn,
	}, nil
}

// for health check of DB
func (d *Database) Ping(ctx context.Context) error {
	return d.Client.DB.PingContext(ctx)
}
