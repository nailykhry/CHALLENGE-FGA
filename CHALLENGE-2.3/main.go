package main

import (
	routers "CHALLENGE-2.3/routes"
)

func main() {
	var PORT = ":8080"
	routers.StartServer().Run(PORT)
}
