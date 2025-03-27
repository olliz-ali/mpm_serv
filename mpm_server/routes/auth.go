package routes

import (
	"mpm_server/mpm_server/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/login", controllers.Login)
	r.GET("/signup", controllers.GetSignup)
	r.POST("/signup", controllers.Signup)
	r.GET("/home", controllers.GetHome)
	r.GET("/premium", controllers.Premium)
	r.GET("/logout", controllers.Logout)
}
