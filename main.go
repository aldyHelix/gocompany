package main

import (
	. "gocompany/database"

	router "gocompany/routes"
)

func main() {
	//ConnectDatabase("root:root@tcp(localhost:3306)/jwt_demo?parseTime=true", "mysql")
	ConnectDatabase(Parameters{ConnectionString: "root:@tcp(localhost:3306)/company?charset=utf8mb4&parseTime=True&loc=Local", ConnectionType: "mysql"})

	port := ":8080"
	Migrate()
	router.NewRoutes().Run(port)
}
