package http

import (
	"net/http"
)

type Config struct {
	Addr string
	HttpHandler http.Handler
}

func Start(config Config) {
	println("Http server start");
	http.ListenAndServe(config.Addr, config.HttpHandler);
};
