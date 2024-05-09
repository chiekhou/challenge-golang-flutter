package routes

import (
	"example/hello/api/controllers/auth"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/signuo", auth.Signup)
}
