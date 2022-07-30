package main

import (
	"mvcGolang/bootstrap"
	"mvcGolang/routes"
)

func main() {
	//gin.SetMode(gin.ReleaseMode) // for release project
	var db = bootstrap.Start()
	//bootstrap.MigrateDatabase() // for migrate and seed database
	var route = routes.Api(db)

	route.Run("localhost:5120")
}
