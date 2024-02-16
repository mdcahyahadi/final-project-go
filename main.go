package main

import (
	"final_project/controllers"
	"final_project/database"
	"final_project/router"
)

func main() {
	db := database.InitDB()

	userRepo := controllers.UserRepo{DB: db}
	photoRepo := controllers.PhotoRepo{DB: db}
	commentRepo := controllers.CommentRepo{DB: db}
	socialMediaRepo := controllers.SocialMediaRepo{DB: db}

	r := router.StartApp(userRepo, photoRepo, commentRepo, socialMediaRepo)
	r.Run(":8080")
}
