package voyage

import (
	"github.com/gin-gonic/gin"
)

// @Summary Invitation groupe de voyage
// @Description Envoie un mail d'invitation afin de de rejoindre un groupen de voyage
// @Tags Voyage
// @Accept json
// @Produce json
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} gin.H "Invitation envoyée"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 404 {object} gin.H "Bad request"
// @Failure 409 {object} gin.H "Conflict"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /send_invitation [post]
func SendInvitation(c *gin.Context) {
	//var emailRequest models.EmailRequest
	/*	err := c.ShouldBindJSON()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}*/

}

// @Summary Rejoindre un groupe de voyage
// @Description Envoie un mail d'invitation afin de de rejoindre un groupen de voyage
// @Tags Voyage
// @Accept json
// @Produce json
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} gin.H "Invitation envoyée"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 404 {object} gin.H "Bad request"
// @Failure 409 {object} gin.H "Conflict"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /send_invitation [post]
func JoinGroupr(c *gin.Context) {

}
