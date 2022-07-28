package main

import (
	"mvcGolang/bootstrap"
	"mvcGolang/routes"
)

func main() {
	var db = bootstrap.Start()
	var route = routes.Api(db)

	route.Run("localhost:5120")
}
