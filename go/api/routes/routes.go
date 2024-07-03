package routes

import (
	"example/hello/api/controllers/activity"
	"example/hello/api/controllers/auth"
    flipping "example/hello/api/controllers/flipping"
	"example/hello/api/controllers/destinations"
	voyage "example/hello/api/controllers/voyages"
	"example/hello/api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/Signup", auth.Signup)
	r.POST("/login", auth.Login)
	r.POST("/logout", auth.Logout)
	r.GET("/profile", middlewares.CheckAuth, auth.UserProfile)
	r.POST("/forgotten_password", auth.MailRecovery)
	r.PUT("/reset_password", auth.ResetPassword)
	r.POST("/join_group", middlewares.CheckAuth, voyage.JoinGroup)
}

func VoyageRoutes(r *gin.Engine) {
	r.GET("/api/voyages", voyage.GetVoyages)
	r.GET("/api/voyages/:id", voyage.GetVoyage)
	r.POST("/api/voyages", voyage.CreateVoyage)
	r.PUT("/api/voyages", voyage.UpdatePutVoyage)
	r.PUT("/api/voyages/hotel", voyage.UpdatePutVoyageHotel)
	r.DELETE("/api/voyages/delete/:id", voyage.DeleteVoyage)
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
}

func FlippingRoutes(r *gin.Engine) {
	r.GET("/api/flipping/feature", flipping.GetFeatureToggle)
	r.PUT("/api/flipping/feature", flipping.UpdateFeatureToggle)

}

