package main

import (
	"goapi/database"
	"goapi/routes"
)

func main() {
	database.ConnectToDatabase()
	routes.HandleRequests()
}
