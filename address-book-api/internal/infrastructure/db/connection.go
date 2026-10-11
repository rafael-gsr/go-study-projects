// Package db contains all database connections and configurations
package db

import (
	"context"
	"fmt"

	"address-book-api/internal/infrastructure/env"

	"github.com/jackc/pgx/v5"
)

func createConnectionString(username string, password string, host string, port int, dbName string, options string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=verify-full&%s", username, password, host, port, dbName, options)
}

func Connect() (*pgx.Conn, error) {
	connectionString := createConnectionString(env.DB_USERNAME, env.DB_PASSWORD, env.DB_HOST, env.DB_PORT, env.DB_NAME, "")
	connection, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}

	return connection, nil
}
