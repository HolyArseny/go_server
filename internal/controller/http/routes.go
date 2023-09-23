package router

import (
	"encoding/json"
	"net/http"
	api "server_example/internal/controller/http/api"
	dto "server_example/internal/dtos"

	"github.com/gorilla/mux"
);

type RouteScheme struct {
	Url  string
	Handler func(http.ResponseWriter, *http.Request)
};

var Routes = map[string][]RouteScheme {
	"GET": {
		RouteScheme {
			Url: "/",
			Handler: func (res http.ResponseWriter, req *http.Request) {
				api.GetRoute1();
			},
		},
		RouteScheme {
			Url: "/users",
			Handler: func (res http.ResponseWriter, req *http.Request) {
				api.GetRoute2();
			},
		},
		RouteScheme {
			Url: "/user/{id}",
			Handler: func (res http.ResponseWriter, req *http.Request) {
				vars := mux.Vars(req);
				id := vars["id"];
				api.GetRouteParam2(id);
			},
		},
	},
	"POST": {
		RouteScheme {
			Url: "/",
			Handler: func (res http.ResponseWriter, req *http.Request) {
				api.PostRoute1();
			},
		},
		RouteScheme {
			Url: "/user",
			Handler: func (res http.ResponseWriter, req *http.Request) {
				var user dto.User

				err := json.NewDecoder(req.Body).Decode(&user)

				if err != nil {
					println(req.RequestURI, "Json decode error");
				}

				api.PostRoute2(user);
			},
		},
	},
};