package api

import "server_example/internal/controller/database"

func GetRoute1() {
	database.DataBase.read("select * from challenges;")
	println("Get route 1")
}

func PostRoute1() {
	println("POST route 1")
}
