package router

import (
	"net/http"

	"github.com/gorilla/mux"
);

func BuildRouter () http.Handler {
	router := mux.NewRouter();

	for method, route := range Routes {
		for _, handler := range route {
			router.HandleFunc(handler.Url, handler.Handler).Methods(method);
		}
	}

	return router;
};
