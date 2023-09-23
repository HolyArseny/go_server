package api

import (
	dto "server_example/internal/dtos"
)

func GetRoute2 () {
	println("Get route 2")
}

func GetRouteParam2 (userId string) {
	println("Get route param 2 user id: ", userId)
}

func PostRoute2 (user dto.User) {
	println("POST route 2");
	println("Username: ", user.Name, "Password: ", user.Password);
}

