package main

import (
	"PROJECT-2/database"
	"PROJECT-2/routers"
)

func main() {
	database.StartDB()
	var PORT = ":8080"
	routers.StartServer().Run(PORT)
}
