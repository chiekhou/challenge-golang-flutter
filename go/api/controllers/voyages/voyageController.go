package voyages

import (
	//"example/hello/api/controllers/requests"
	//"example/hello/bin/utils"
	"example/hello/internal/initializers"
	"example/hello/internal/models"
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
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer Add access token here)
// @Success      200  {object}  models.Voyage
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
//
// @Router	 	/api/voyages [get]
func GetVoyages(c *gin.Context) {
	currentUser, exist := c.Get("currentUser")
	if !exist {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Not authorized"})
		return
	}

	user := currentUser.(models.User)

	var voyages []models.Voyage
	result := initializers.DB.Where("user_id = ?", user.ID).Preload("Activities").Preload("Hotels").Find(&voyages)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, DestinationResponse{Data: voyages})

}

// ShowVoyage godoc
// @Summary      Show a voyage
// @Description  get string by ID
// @Tags         Voyages
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Voyages ID"
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer Add access token here)
// @Success      200  {object}  models.Voyage
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/voyages/{id} [get]
func GetVoyage(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	currentUser, exist := c.Get("currentUser")
	if !exist {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Not authorized"})
		return
	}

	user := currentUser.(models.User)

	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID format"})
		return
	}

	var voyage models.Voyage
	result := initializers.DB.Preload("Activities").Preload("Hotels").First(&voyage, id)

	if voyage.UserId != user.ID {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Non Authorized User"})
		return
	}

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
// @Summary     Add a voyage
// @Description Add by JSON voyage
// @Tags        Voyages
// @Accept      json
// @Produce     json
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer Add access token here)
// @Param       voyage body models.Voyage true "Add voyage"
// @Success     200 {object} models.Voyage
// @Failure     400 {object} ErrorResponse
// @Failure     404 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /api/voyages [post]
func CreateVoyage(c *gin.Context) {
	user, exist := c.Get("currentUser")
	if !exist {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Not authorized"})
		return
	}

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
		UserId:      user.(models.User).ID,
	}

	if err := initializers.DB.Create(&voyage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("Erreur de création de voyage:", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": voyage})
	fmt.Println("Voyage créé avec succès:", voyage)

}

// updateVoyage with Put godoc
// @Summary Update a trip
// @Description Update a trip by ID
// @Tags Voyages
// @Accept json
// @Produce json
// @Param voyage body models.Voyage true "Voyage data"
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer Add access token here)
// @Success      200  {object}  models.Voyage
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router /api/voyages [put]
func UpdatePutVoyage(c *gin.Context) {
	currentUser, exist := c.Get("currentUser")
	if !exist {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Not authorized"})
		return
	}

	user, ok := currentUser.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Unable to retrieve user information"})
		return
	}

	var body models.Voyage
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if body.UserId != user.ID {
		c.JSON(http.StatusForbidden, ErrorResponse{Error: "You do not have permission to update this voyage"})
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
	voyage.UserId = user.ID

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
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer Add access token here)
// @Success      200  {object}  models.Voyage
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router /api/voyages/hotel [put]
func UpdatePutVoyageHotel(c *gin.Context) {
	currentUser, exist := c.Get("currentUser")
	if !exist {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Not authorized"})
		return
	}
	user, ok := currentUser.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Unable to retrieve user information"})
	}

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

	if voyage.UserId != user.ID {
		c.JSON(http.StatusForbidden, ErrorResponse{Error: "You do not have permission to update this voyage"})
		return
	}

	voyage.Destination = body.Destination
	voyage.DateAller = body.DateAller
	voyage.DateRetour = body.DateRetour
	voyage.Activities = body.Activities
	voyage.Hotels = body.Hotels
	voyage.UserId = user.ID

	if err := initializers.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&voyage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, voyage)
}

// DeleteVoyage godoc
//
//	@Summary		Delete a voyage
//	@Description	Delete by voyage ID
//	@Tags			Voyages
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Voyage ID"	Format(int64)
//
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer Add access token here)
// @Success      200  {object}  models.Voyage
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
//
//	@Router			/api/voyages/delete/{id} [delete]
func DeleteVoyage(c *gin.Context) {
	currentUser, exist := c.Get("currentUser")
	if !exist {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Not authorized"})
		return
	}
	user, ok := currentUser.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Unable to retrieve user information"})
	}

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

	if voyage.UserId != user.ID {
		c.JSON(http.StatusForbidden, ErrorResponse{Error: "You do not have permission to delete this voyage"})
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
