package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

type db struct {
	conn
}

func (d *db) create(query) {
	result, err := Create(d, query)
	if err != nil {
		return false, err
	}
	return true, result
}

func (d *db) read(query) {
	result, err := Read(d, query)
	if err != nil {
		return false, err
	}
	return true, result

}

func (d *db) update(query) {
	result, err := Update(d, query)
	if err != nil {
		return false, err
	}
	return true, result
}

func (d *db) delete(query) {
	result, err := Delete(d, query)
	if err != nil {
		return false, err
	}
	return true, result
}

func (d db) closeConnection() {
	d.conn.Close(context.Background())
}

var DataBase *db

func InitDB() {
	if DataBase != nil {
		return
	}

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	DataBase := db{conn}

	defer DataBase.closeConnection()
}
