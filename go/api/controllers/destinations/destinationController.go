package destinations

import (
	"errors"
	"example/hello/internal/initializers"
	"example/hello/internal/models"
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

// List Destinations godoc
//
//	@Summary		List Destinations
//	@Description	Get Destinations
//	@Tags			Destinations
//	@Accept			json
//	@Produce		json
//
// @Success      200  {object}  models.Destination
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
//
// @Router			/api/destinations [get]
func GetDestinations(c *gin.Context) {
	var destinations []models.Destination
	result := initializers.DB.Preload("Activities").Preload("Hotels").Find(&destinations)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, DestinationResponse{Data: destinations})
}

// ShowDestination godoc
//
// @Summary      Show a destination
// @Description  get string by ID
// @Tags         Destinations
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Destination ID"
// @Success      200  {object}  models.Destination
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/destinations/{id} [get]
func GetDestination(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID format"})
		return
	}

	var destination models.Destination
	result := initializers.DB.Preload("Activities").Preload("Hotels").First(&destination, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, ErrorResponse{Error: "Destination not found"})
		} else {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: result.Error.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, destination)
}

// AddDestination godoc
// @Summary     Add a destination
// @Description Add by JSON destination
// @Tags        Destinations
// @Accept      json
// @Produce     json
// @Param       destination body models.Destination true "Add destination"
// @Success     200 {object} models.Destination
// @Failure     400 {object} ErrorResponse
// @Failure     404 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer Add access token here)
//
// @Router      /api/destinations [post]
func CreateDestination(c *gin.Context) {
	var input struct {
		Departure time.Time `json:"departure"`
		Return    time.Time `json:"return_date"`
		Name      string    `json:"name"`
		Image     string    `json:"image"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	destination := models.Destination{
		Name:  input.Name,
		Image: input.Image,
	}

	if err := initializers.DB.Create(&destination).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, DestinationResponse{Data: destination})
}

// UpdateDestination godoc
// @Summary		Update a destination
// @Description	Update by json destination
// @Tags		Destinations
// @Accept		json
// @Produce		json
// @Param       id path int true "Destination ID"
// @Param       destination body models.Destination true "Update Destination"
// @Success      200  {object}  models.Destination
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router			/api/destinations/update/{id} [patch]
func UpdateDestination(c *gin.Context) {
	var input struct {
		Name       string            `json:"name"`
		Image      string            `json:"image"`
		UserID     *uint             `json:"user_id"`
		Activities []models.Activity `json:"activities"`
		Hotels     []models.Hotel    `json:"hotels"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	var destination models.Destination
	if err := initializers.DB.Preload("Activities").Preload("Hotels").First(&destination, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Destination not found"})
		return
	}

	var activities []models.Activity
	if len(input.Activities) > 0 {
		initializers.DB.Where("id IN ?", input.Activities).Find(&activities)
	}

	var hotels []models.Hotel
	if len(input.Hotels) > 0 {
		initializers.DB.Where("id IN ?", input.Hotels).Find(&hotels)
	}

	destination.Name = input.Name
	destination.Image = input.Image
	destination.Activities = input.Activities
	destination.Hotels = input.Hotels

	if err := initializers.DB.Save(&destination).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, DestinationResponse{Data: destination})
}

// DeleteDestination godoc
//
//	@Summary		Delete a destination
//	@Description	Delete by destination ID
//	@Tags			Destinations
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Destination ID"	Format(int64)
//
// @Success      200  {object}  models.Destination
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer Add access token here)
//
//	@Router			/api/destinations/delete/{id} [delete]
func DeleteDestination(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID format"})
		return
	}

	var destination models.Destination
	if err := initializers.DB.First(&destination, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Destination not found"})
		return
	}

	// Supprimer les activités associées
	if err := initializers.DB.Model(&destination).Association("Activities").Clear(); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	// Supprimer les hôtels associées
	if err := initializers.DB.Model(&destination).Association("Hotels").Clear(); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	// Supprimer la destination
	if err := initializers.DB.Delete(&destination).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Data: true})
}

// AddActivityToDestination godoc
//
//	@Summary		Add a activity to destination
//	@Description	add by json activity destination
//	@Tags			Destinations
//	@Accept			json
//	@Produce		json
//
// @Param        id path int true "Destination ID"
// @Param        activity       body models.Activity true "Activity to add"
// @Success      200  {object}  models.Destination
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
//
// @Router			/api/destinations/{id}/activity [post]
func CreateActivityDestination(c *gin.Context) {
	destinationID := c.Param("id")
	var activity models.DestinationActivity

	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	var destination models.Destination
	if err := initializers.DB.First(&destination, destinationID).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Destination not found"})
		return
	}

	activity.DestinationID = destination.ID

	if err := initializers.DB.Create(&activity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	if err := initializers.DB.Model(&destination).Association("Activities").Append(&activity); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, DestinationResponse{Data: destination})
}

// VerifyActivtyName godoc
//
//	@Summary		Verify an activity Name
//	@Description	verify by json activity name
//	@Tags			Destinations
//	@Accept			json
//	@Produce		json
//
// @Param        id path int true "Destination ID"
// @Param        name path string true "Activity Name"
// @Success      200  {object}  models.Destination
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
//
// @Router			/api/destination/{id}/activities/verify/{name} [get]
func VerifyActivtyName(c *gin.Context) {
	destinationID := c.Param("id")
	activityName := c.Param("name")

	var destination models.Destination
	if err := initializers.DB.First(&destination, destinationID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Destination not found"})
		return
	}

	// Rechercher l'ID de la destination associée à l'activité avec le nom donné
	var destinationActivity models.DestinationActivity
	if err := initializers.DB.Where("destination_id = ? AND activity_id IN (SELECT id FROM activities WHERE name = ?)", destinationID, activityName).First(&destinationActivity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Aucune activité n'a été trouvée avec ce nom
			c.JSON(http.StatusNotFound, gin.H{"error": "No activity found with this name"})
			return
		}
		// Autre erreur
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// L'activité a été trouvée avec succès
	c.JSON(http.StatusOK, gin.H{"destinationActivity": destinationActivity})
}
