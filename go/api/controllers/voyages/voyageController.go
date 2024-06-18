package voyages

import (
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
// @Success      200  {object}  models.Voyage
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
//
// @Router	 	/api/voyages [get]
func GetVoyages(c *gin.Context) {

	var voyages []models.Voyage
	result := initializers.DB.Preload("Activities").Find(&voyages)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, DestinationResponse{Data: voyages})

}

// ShowVoyage godoc
// @Summary      Show a groupeVoyage
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
	result := initializers.DB.Preload("Activities").First(&voyage, id)

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

	var input struct {
		Destination string            `json:"destination"`
		Date        time.Time         `json:"date"`
		Activities  []models.Activity `json:"activities"` // List of activity IDs
	}

	// Bind JSON input to the input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("Erreur de binding JSON:", err.Error())
		return
	}

	var activities []models.Activity
	if len(input.Activities) > 0 {
		initializers.DB.Where("id IN ?", input.Activities).Find(&activities)
	}

	voyage := models.Voyage{
		Destination: input.Destination,
		Date:        input.Date,
		Activities:  input.Activities,
	}

	if err := initializers.DB.Create(&voyage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("Erreur de création de groupeVoyage:", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": voyage})
	fmt.Println("Voyage créé avec succès:", voyage)

}

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
	if err := initializers.DB.Preload("Activities").First(&voyage, body.ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Trip not found"})
		return
	}

	voyage.Destination = body.Destination
	voyage.Date = body.Date
	voyage.Activities = body.Activities

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

	// Supprimer la destination
	if err := initializers.DB.Delete(&voyage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Data: true})
}
