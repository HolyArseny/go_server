package api

import "context"

func Read(conn *pgx.Conn, query string) (string, error) {
	var result string
	err := conn.QueryRow(context.Background(), query).Scan(&result)

	if err != nil {
		return "error", err
	}

	return result, nil
}
