package api

import (
	database "server_example/internal/controller/database"
)

func GetRoute1() {
	database.DataBase.Read("select name from users where id = 1;")
	println("Get route 1")
}

func PostRoute1() {
	println("POST route 1")
}
