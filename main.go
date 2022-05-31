package main

import (
	"gocompany/models"

	router "gocompany/routes"
)

func main() {
	models.ConnectDatabase()

	router.NewRoutes().Run()
}
