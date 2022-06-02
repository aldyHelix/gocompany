package main

import (
	. "gocompany/database"
	. "gocompany/env"
	router "gocompany/routes"
)

func main() {
	ConnectDatabase(Parameters{ConnectionString: "root:@tcp(localhost:3306)/company?charset=utf8mb4&parseTime=True&loc=Local", ConnectionType: "mysql"})

	port := ":" + Env("PORT") //Port from env
	Migrate()
	router.NewRoutes().Run(port)
}
