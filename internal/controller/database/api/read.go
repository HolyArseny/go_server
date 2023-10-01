package api

import "context"

func Read(conn *pgx.Conn, query string) (string, bool) {
	var result string
	err := conn.QueryRow(context.Background(), query).Scan(&result)

	if err != nil {
		return "error", true
	}

	return result, nil
}
