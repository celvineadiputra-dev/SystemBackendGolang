package main

import (
	"mvcGolang/bootstrap"
	"mvcGolang/routes"
)

func main() {
	var db = bootstrap.Start()
	bootstrap.MigrateDatabase()
	var route = routes.Api(db)

	route.Run("localhost:5120")
}
