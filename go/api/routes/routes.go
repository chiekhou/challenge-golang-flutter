package routes

import (
	"example/hello/api/controllers/auth"
	"example/hello/api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/Signup", auth.Signup)
	r.POST("/login", auth.Login)
	//r.POST("/logout", auth.Logout)
	r.GET("/profile", middlewares.CheckAuth, auth.UserProfile)
}
