package main

import (
	"server_example/internal/config"
	router "server_example/internal/controller/http"
	http "server_example/internal/server"
)

func main() {
	println("App start");

	httpConfig := http.Config {
		Addr: config.GetEnv().HTTP_PORT,
		HttpHandler: router.BuildRouter(),
	};

	http.Start(httpConfig);
}