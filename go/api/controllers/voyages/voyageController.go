package voyages

import (
	"example/hello/api/controllers/requests"
	"example/hello/bin/utils"
	"example/hello/internal/initializers"
	"example/hello/internal/models"
	mailer2 "example/hello/pkg/mailer"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type DestinationResponse struct {
	Data interface{} `json:"data"`
}

type SuccessResponse struct {
	Data bool `json:"data"`
}

// List Voyages godoc
// @Summary		List Voyages
// @Description	Get Voyages
// @Tags			Voyages
// @Accept			json
// @Produce		json
// @Success      200  {object}  models.Voyage
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
//
// @Router	 	/api/voyages [get]
func GetVoyages(c *gin.Context) {

	var voyages []models.Voyage
	result := initializers.DB.Preload("Activities").Preload("Hotels").Find(&voyages)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, DestinationResponse{Data: voyages})

}

// ShowVoyage godoc
<<<<<<< HEAD
// @Summary      Show a voyage
=======
// @Summary      Show a groupeVoyage
>>>>>>> origin/feature/merge_voyage
// @Description  get string by ID
// @Tags         Voyages
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Voyages ID"
// @Success      200  {object}  models.Voyage
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/voyages/{id} [get]
func GetVoyage(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID format"})
		return
	}

	var voyage models.Voyage
	result := initializers.DB.Preload("Activities").Preload("Hotels").First(&voyage, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, ErrorResponse{Error: "Destination not found"})
		} else {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: result.Error.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, voyage)
}

// AddVoyage godoc
// @Summary     Add a groupeVoyage
// @Description Add by JSON groupeVoyage
// @Tags        Voyages
// @Accept      json
// @Produce     json
// @Param       groupeVoyage body models.Voyage true "Add groupeVoyage"
// @Success     200 {object} models.Voyage
// @Failure     400 {object} ErrorResponse
// @Failure     404 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /api/voyages [post]
func CreateVoyage(c *gin.Context) {

	var featureToggles = map[string]bool{
		"active_voyage": true,
	}
	enabled, exists := featureToggles["active_voyage"]
	if !exists || !enabled {
		c.JSON(http.StatusForbidden, gin.H{"error": "Vous ne pouvez pas créer un voyage"})
		return
	}

	var input struct {
		Destination string            `json:"destination"`
		DateAller   time.Time         `json:"dateAller"`
		DateRetour  time.Time         `json:"dateRetour"`
		Activities  []models.Activity `json:"activities"`
		Hotels      []models.Hotel    `json:"hotels"`
	}

	// Bind JSON input to the input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("Erreur de binding JSON:", err.Error())
		return
	}

	voyage := models.Voyage{
		Destination: input.Destination,
		DateAller:   input.DateAller,
		DateRetour:  input.DateRetour,
		Activities:  input.Activities,
		Hotels:      input.Hotels,
	}

	if err := initializers.DB.Create(&voyage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("Erreur de création de groupeVoyage:", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": voyage})
	fmt.Println("Voyage créé avec succès:", voyage)

}

<<<<<<< HEAD
=======
// UpdateVoyage godoc
// @Summary		Update a groupeVoyage
// @Description	Update by json destination
// @Tags		Voyages
// @Accept		json
// @Produce		json
// @Param       id path int true "Voyage ID"
// @Param       groupeVoyage body models.Voyage true "Update Voyage"
// @Success      200  {object}  models.Voyage
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router			/api/voyages/update/{id} [patch]
func UpdateVoyage(c *gin.Context) {
	var input struct {
		Destination string            `json:"destination"`
		Date        time.Time         `json:"date"`
		Activities  []models.Activity `json:"activities"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var voyage models.Voyage
	if err := initializers.DB.Preload("Activities").First(&voyage, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Voyage not found"})
		return
	}

	var activities []models.Activity
	if len(input.Activities) > 0 {
		initializers.DB.Where("id IN ?", input.Activities).Find(&activities)
	}

	voyage.Destination = input.Destination
	voyage.Date = input.Date
	voyage.Activities = input.Activities

	if err := initializers.DB.Save(&voyage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": voyage})
}

>>>>>>> origin/feature/merge_voyage
// updateVoyage with Put godoc
// @Summary Update a trip
// @Description Update a trip by ID
// @Tags Voyages
// @Accept json
// @Produce json
// @Param groupeVoyage body models.Voyage true "Voyage data"
// @Success      200  {object}  models.Voyage
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router /api/voyages [put]
func UpdatePutVoyage(c *gin.Context) {
	var body models.Voyage
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var voyage models.Voyage
	if err := initializers.DB.Preload("Activities").Preload("Hotels").First(&voyage, body.ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Voyage not found"})
		return
	}

	voyage.Destination = body.Destination
	voyage.DateAller = body.DateAller
	voyage.DateRetour = body.DateRetour
	voyage.Activities = body.Activities
	//voyage.Hotels = body.Hotels

	if err := initializers.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&voyage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, voyage)
}

// updateVoyageHotel with Put godoc
// @Summary Update a trip
// @Description Update a trip by ID
// @Tags Voyages
// @Accept json
// @Produce json
// @Param voyage body models.Voyage true "Voyage data"
// @Success      200  {object}  models.Voyage
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router /api/voyages/hotel [put]
func UpdatePutVoyageHotel(c *gin.Context) {
	var body models.Voyage
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var voyage models.Voyage
	if err := initializers.DB.Preload("Activities").Preload("Hotels").First(&voyage, body.ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Voyage not found"})
		return
	}

	voyage.Destination = body.Destination
	voyage.DateAller = body.DateAller
	voyage.DateRetour = body.DateRetour
	voyage.Activities = body.Activities
	voyage.Hotels = body.Hotels

	if err := initializers.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&voyage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, voyage)
}

// DeleteVoyage godoc
//
//	@Summary		Delete a groupeVoyage
//	@Description	Delete by groupeVoyage ID
//	@Tags			Voyages
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Voyage ID"	Format(int64)
//
// @Success      200  {object}  models.Voyage
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
//
//	@Router			/api/voyages/delete/{id} [delete]
func DeleteVoyage(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID format"})
		return
	}

	var voyage models.Voyage
	if err := initializers.DB.First(&voyage, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Destination not found"})
		return
	}

	// Supprimer les activités associées
	if err := initializers.DB.Model(&voyage).Association("Activities").Clear(); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	if err := initializers.DB.Model(&voyage).Association("Hotels").Clear(); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	// Supprimer la destination
	if err := initializers.DB.Delete(&voyage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Data: true})
}

// @Summary Invitation groupe de voyage
// @Description Envoie un mail d'invitation afin de de rejoindre un groupen de voyage
// @Tags Voyages
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
	var emailRequest requests.InvitationGroupRequest

	// Chercher si le groupe de voyage existe
	var group models.GroupeVoyage
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&group).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Groupe non trouvé"})
		return
	}

	if err := c.ShouldBindJSON(&emailRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(emailRequest.Email, group.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	type emailData struct {
		UserFound models.User
		Token     string
		GroupID   uint
	}
	var email emailData
	email.Token = token
	email.GroupID = group.ID

	initializers.DB.Where("email = ?", emailRequest.Email).First(&userFound)
	email.UserFound = userFound

	if userFound.ID == 0 {
		mailer2.SendGoMail(emailRequest.Email,
			"Inscription",
			"./pkg/mailer/templates/registry.html",
			email)
	} else {
		mailer2.SendGoMail(emailRequest.Email,
			"Invitation dans un groupe de voyage",
			"./pkg/mailer/templates/invite.html",
			email)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Email envoyé",
		"token":   token,
	})
}

// @Summary Rejoindre un groupe de voyage
// @Description Envoie un mail d'invitation afin de de rejoindre un groupen de voyage
// @Tags Voyages
// @Accept json
// @Produce json
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} gin.H "Invitation envoyée"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 404 {object} gin.H "Bad request"
// @Failure 409 {object} gin.H "Conflict"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /join_group [post]
func JoinGroup(c *gin.Context) {
	var request struct {
		Token string `json:"token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email, groupID, err := utils.ParseToken(request.Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token invalide"})
		return
	}

	var user models.User
	if err := initializers.DB.Where("email = ?", email).First(&user).Error; err != nil {
		// Si l'utilisateur n'existe pas, le créer
		user = models.User{Email: email}
		if err := initializers.DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création de l'utilisateur"})
			return
		}
	}

	var group models.GroupeVoyage
	if err := initializers.DB.Where("id = ?", groupID).First(&group).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Groupe non trouvé"})
		return
	}

	// Ajouter l'utilisateur au groupe
	if err := initializers.DB.Model(&group).Association("Members").Append(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de l'ajout de l'utilisateur au groupe"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Groupe rejoint"})
}
