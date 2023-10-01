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

func (d db) Read(query string) {
	result, err := dbApi.Read(d.connector, query)
	if err != nil {
		fmt.Println("read query eror: ", err)
	}
	fmt.Println("read query result", result)
}

func (d db) Create(query string) {
	result, err := dbApi.Create(d.connector, query)
	if err != nil {
		fmt.Println("create query error: ", err)
	}
	fmt.Println("create query result", result)
}

func (d db) Update(query string) {
	result, err := dbApi.Update(d.connector, query)
	if err != nil {
		fmt.Println("update query error: ", err)
	}
	fmt.Println("update query result: ", result)
}

func (d db) Delete(query string) {
	result, err := dbApi.Delete(d.connector, query)
	if err != nil {
		fmt.Println("delet query error: ", err)
	}
	fmt.Println("delete query result: ", result)
}

func (d db) CloseConnection() {
	println("DB CLOSE")
	d.connector.Close(context.Background())
}

var DataBase db

func InitDB() {
	println("DB INIT")

	if DataBase.connector != nil {
		return
	}

	urlExample := "postgres://postgres:jncg8azx@localhost:5432/k_learn"
	connector, err := pgx.Connect(context.Background(), urlExample)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	DataBase = db{connector}
}
