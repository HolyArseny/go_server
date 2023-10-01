package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
	dbApi "server_example/internal/controller/database/api"
)

type db struct {
	connector *pgx.Conn
}

// func (d db) create(query string) {
// 	result, err := dbApi.Create(d.connector, query)
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, result
// }

func (d db) read(query string) {
	result, err := dbApi.Read(d.connector, query)
	if err != nil {
		return false, err
	}
	return true, result
}

// func (d db) update(query string) {
// 	result, err := dbApi.Update(d.connector, query)
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, result
// }

// func (d db) delete(query string) {
// 	result, err := dbApi.Delete(d.connector, query)
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, result
// }

func (d db) closeConnection() {
	d.connector.Close(context.Background())
}

var DataBase *db

func InitDB() {
	if DataBase != nil {
		return
	}

	urlExample := "postgres://postgres:jncg8azx@localhost:5432/k_learn"
	connector, err := pgx.Connect(context.Background(), urlExample)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	DataBase := db{connector}

	defer DataBase.closeConnection()
}
