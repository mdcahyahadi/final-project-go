package main

import (
	"final_project/controllers"
	"final_project/database"
	"final_project/router"
)

func main() {
	db := database.InitDB()

	userRepo := controllers.UserRepo{DB: db}

	r := router.StartApp(userRepo)
	r.Run(":8080")
}
