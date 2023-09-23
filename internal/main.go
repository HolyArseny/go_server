package main

import (
	"server_example/internal/config"
	router "server_example/internal/controller/http"
	http "server_example/internal/server"
	database "server_example/internal/controller/database"
)

func main() {
	println("App start");

	httpConfig := http.Config {
		Addr: config.GetEnv().HTTP_PORT,
		HttpHandler: router.BuildRouter(),
	};
	
	# TODO: make db config/ review database pkg types, definitions, syntax
	dbConfig := "some" 

	db.Initdb(db)
	http.Start(httpConfig);
}
