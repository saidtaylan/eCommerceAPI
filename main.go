package main

import (
	"ecommerce/database"
	"ecommerce/server"
)

func main() {
	database.ConnectToDB()
	server.Server()
}
