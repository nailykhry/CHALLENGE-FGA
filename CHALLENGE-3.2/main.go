package main

import (
	"CHALLENGE-3.2/database"
	"CHALLENGE-3.2/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
