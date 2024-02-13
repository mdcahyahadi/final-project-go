package router

import (
	"final_project/controllers"
	"final_project/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp(c controllers.UserRepo) *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.GET("/", c.GetAllUser)
		userRouter.POST("/register", c.UserRegister)
		userRouter.POST("/login", c.UserLogin)
		userRouter.PUT("/:userId", middlewares.Authentication(), c.UserUpdate)
		userRouter.DELETE("/", middlewares.Authentication(), c.UserDelete)
	}

	return r
}
