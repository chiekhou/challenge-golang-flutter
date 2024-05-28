package routes

import (
	"example/hello/api/controllers/auth"
	"example/hello/api/controllers/voyage"
	"example/hello/api/middlewares"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	store := cookie.NewStore([]byte(os.Getenv("SECRET")))
	r.Use(sessions.Sessions("my_session", store))

	r.POST("/Signup", auth.Signup)
	r.POST("/login", auth.Login)
	r.POST("/logout", auth.Logout)
	r.GET("/profile", middlewares.CheckAuth, auth.UserProfile)
	r.POST("/forgotten_password", auth.MailRecovery)
	r.PUT("/reset_password", auth.ResetPassword)

	r.POST("/groups", middlewares.CheckAuth, voyage.CreateGroup)
	r.POST("/groups/:group_id/join", middlewares.CheckAuth, voyage.JoinGroup)
	r.PUT("/groups/:id/budget", middlewares.CheckAuth, voyage.UpdateBudget)
}
