package router

import (
	"final_project/controllers"
	"final_project/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp(u controllers.UserRepo, p controllers.PhotoRepo, c controllers.CommentRepo, s controllers.SocialMediaRepo) *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.GET("/", u.GetAllUser)
		userRouter.POST("/register", u.UserRegister)
		userRouter.POST("/login", u.UserLogin)
		userRouter.PUT("/:userId", middlewares.Authentication(), u.UserUpdate)
		userRouter.DELETE("/", middlewares.Authentication(), u.UserDelete)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.GET("/", p.GetAllPhoto)
		photoRouter.GET("/:photoId", p.GetPhotoByID)
		photoRouter.POST("/", p.CreatePhoto)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), p.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), p.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.GET("/", c.GetAllComment)
		commentRouter.GET("/:commentId", c.GetCommentByID)
		commentRouter.POST("/", c.CreateComment)
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), c.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), c.DeleteComment)
	}

	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.GET("/", s.GetAllSocialMedia)
		socialMediaRouter.GET("/:socialMediaId", s.GetSocialMediaByID)
		socialMediaRouter.POST("/", s.CreateSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), s.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), s.DeletSocialeMedia)
	}

	return r
}
