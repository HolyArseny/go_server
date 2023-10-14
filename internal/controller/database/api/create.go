package api

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Create(conn *pgxpool.Pool, query string) (string, error) {
	var result string
	err := conn.QueryRow(context.Background(), query).Scan(&result)

	if err != nil {
		return "error", err
	}

	return result, nil
}
