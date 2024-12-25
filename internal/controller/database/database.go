package database

import (
	"context"
	"fmt"
	"os"
	dbApi "server_example/internal/controller/database/api"
	dbModels "server_example/internal/controller/database/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type db struct {
	_pool *pgxpool.Pool
}

func (d db) Read(query string) {
	result, err := dbApi.Read(d._pool, query)
	if err != nil {
		fmt.Println("read query eror: ", err)
	}
	fmt.Println("read query result", result)
}

func (d db) Create(query string) {
	result, err := dbApi.Create(d._pool, query)
	if err != nil {
		fmt.Println("create query error: ", err)
	}
	fmt.Println("create query result", result)
}

func (d db) Update(query string) {
	result, err := dbApi.Update(d._pool, query)
	if err != nil {
		fmt.Println("update query error: ", err)
	}
	fmt.Println("update query result: ", result)
}

func (d db) Delete(query string) {
	result, err := dbApi.Delete(d._pool, query)
	if err != nil {
		fmt.Println("delet query error: ", err)
	}
	fmt.Println("delete query result: ", result)
}

func (d db) CloseConnection() {
	println("DB CLOSE")
	d._pool.Close()
}

func (d db) syncSchema() {
	for _, schema := range dbModels.Schemas {
		fmt.Print("MIGRATING ", schema.Name, "... ")

		_, error := d._pool.Query(context.Background(), schema.Query)

		if error != nil {
			fmt.Println("FAILED: ", error)
			os.Exit(1)
		}

		fmt.Println("SUCCEED")
	}
}

var DataBase db

func InitDB() {
	println("DB INIT")

	if DataBase._pool != nil {
		return
	}

	urlExample := "postgres://postgres:jncg8azx@localhost:5432/k_learn"
	connector, err := pgxpool.New(context.Background(), urlExample)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	DataBase = db{connector}
	DataBase.syncSchema()
}
