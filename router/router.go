package router

import (
	"final-project/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)

	}

	// photoRouter := r.Group("/photo")
	// {
	// 	photoRouter.POST("/register", controllers.CreatePhoto)

	// }

	return r
}
