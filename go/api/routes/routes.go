package routes

import (
	"example/hello/api/controllers/activity"
	"example/hello/api/controllers/auth"
	"example/hello/api/controllers/destinations"
	flipping "example/hello/api/controllers/flipping"
	groupVoyage "example/hello/api/controllers/groupeVoyage"
	"example/hello/api/controllers/sockets"
	"example/hello/api/controllers/user"
	voyage "example/hello/api/controllers/voyages"
	"example/hello/api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/Signup", auth.Signup)
	r.POST("/login", auth.Login)
	r.POST("/logout", middlewares.CheckAuth, auth.Logout)
	r.GET("/profile", middlewares.CheckAuth, auth.UserProfile)
	r.POST("/forgotten_password", auth.MailRecovery)
	r.PUT("/reset_password", middlewares.CheckAuth, auth.ResetPassword)
}

func VoyageRoutes(r *gin.Engine) {
	r.GET("/api/voyages", middlewares.CheckAuth, voyage.GetVoyages)
	r.GET("/api/voyages/:id", middlewares.CheckAuth, voyage.GetVoyage)
	r.POST("/api/voyages", middlewares.CheckAuth, voyage.CreateVoyage)
	r.PUT("/api/voyages", middlewares.CheckAuth, voyage.UpdatePutVoyage)
	r.PUT("/api/voyages/hotel", middlewares.CheckAuth, voyage.UpdatePutVoyageHotel)
	r.DELETE("/api/voyages/delete/:id", middlewares.CheckAuth, voyage.DeleteVoyage)
}

func DestinationRoutes(r *gin.Engine) {
	r.GET("/api/destinations", destinations.GetDestinations)
	r.GET("/api/destinations/:id", destinations.GetDestination)
	r.GET("/api/destination/:id/activities/verify/:name", destinations.VerifyActivtyName)
	r.POST("/api/destinations", destinations.CreateDestination)
	r.POST("/api/destinations/:id/activity", destinations.CreateActivityDestination)
	r.PATCH("/api/destinations/update/:id", destinations.UpdateDestination)
	r.DELETE("/api/destinations/delete/:id", destinations.DeleteDestination)

}

func ActivityRoutes(r *gin.Engine) {
	r.POST("/api/activity/images", activity.UploadImage)
	r.POST("/create_group", middlewares.CheckAuth, groupVoyage.CreateGroup)
	r.DELETE("/groupes/:group_id/delete_group", middlewares.CheckAuth, groupVoyage.DeleteGroup)
	r.GET("/groupes", groupVoyage.GetAllGroups)
	r.GET("/groupes/:group_id/join", groupVoyage.Join)
	r.POST("/groupes/:group_id/send_invitation", middlewares.CheckAuth, groupVoyage.SendInvitation)
	r.PUT("/groupes/:group_id/update_budget", middlewares.CheckAuth, groupVoyage.UpdateBudget)
	r.GET("/groupes/my_groups", middlewares.CheckAuth, groupVoyage.GetMyGroups)
	r.GET("/groupes/:group_id", middlewares.CheckAuth, groupVoyage.GetGroupById)
	r.DELETE("/groupes/:group_id/member/:member_id/delete_member", middlewares.CheckAuth, groupVoyage.DeleteGroupMember)
}

func FlippingRoutes(r *gin.Engine) {
	r.GET("/api/flipping/feature", flipping.GetFeatureToggle)
	r.PUT("/api/flipping/feature", flipping.UpdateFeatureToggle)

}

func UsersRoutes(r *gin.Engine) {
	r.POST("/api/users", middlewares.CheckAuth, user.CreateUser)
	r.GET("/api/users", middlewares.CheckAuth, user.GetUsers)
	r.GET("/api/users/:id", middlewares.CheckAuth, user.GetUser)
	r.PUT("/api/users/:id", middlewares.CheckAuth, user.UpdateUser)
	r.DELETE("/api/users/:id", middlewares.CheckAuth, user.DeleteUser)

}

func SocketRoutes(r *gin.Engine) {
	r.GET("/api/messages/:groupe_voyage_id", sockets.GetPreviousMessages)
}
