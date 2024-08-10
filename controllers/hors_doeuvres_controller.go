package controllers

import (
	"log"
	"net/http"

	"github.com/ax-vasquez/wedding-site-api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HorsDoeuvresData struct {
	HorsDoeuvres []models.HorsDoeuvres `json:"hors_doeuvres"`
}

type V1_API_RESPONSE_HORS_DOEUVRES struct {
	V1_API_RESPONSE
	Data HorsDoeuvresData `json:"data"`
}

// GetHorsDoeuvres gets one or all hors doeuvres
//
//	@Summary      creates an hors doeuvres
//	@Description  Gets the selected hors doeuvres for the given user ID (empty array if no selection has been made), or a list of all available entrees
//	@Tags         hors doeuvres
//	@Accept       json
//	@Produce      json
//	@Success      200  {object}  V1_API_RESPONSE_HORS_DOEUVRES
//	@Failure      500  {object}  V1_API_RESPONSE_HORS_DOEUVRES
//	@Param 		  user_id  path uuid.UUID true "User ID"
//	@Router       /entree [get]
//	@Router       /user/{user_id}/horsdoeuvres [get]
func GetHorsDoeuvres(c *gin.Context) {
	idStr := c.Param("id")
	var response V1_API_RESPONSE_HORS_DOEUVRES
	var status int
	var horsDoeuvres []models.HorsDoeuvres
	id, err := uuid.Parse(idStr)
	if err == nil {
		status = http.StatusOK
		horsDoeuvres, err = models.FindHorsDoeuvresForUser(id)
		if err != nil {
			status = http.StatusInternalServerError
			log.Println(err.Error())
			response.Message = "Internal server error"
		} else {
			status = http.StatusOK
		}
	} else {
		var err error
		horsDoeuvres, err = models.FindHorsDoeuvres()
		if err != nil {
			status = http.StatusInternalServerError
			log.Println(err.Error())
			response.Message = "Internal server error"
		} else {
			status = http.StatusOK
		}
	}
	response.Status = status
	response.Data = HorsDoeuvresData{
		HorsDoeuvres: horsDoeuvres,
	}
	c.JSON(status, response)
}

// CreateHorsDoeuvres creates an hors doeuvres
//
//	@Summary      creates an hors doeuvres
//	@Description  Creates an hors doeuvres and return the new record's data to the caller
//	@Tags         hors doeuvres
//	@Accept       json
//	@Produce      json
//	@Success      200  {object}  V1_API_RESPONSE_HORS_DOEUVRES
//	@Failure      500  {object}  V1_API_RESPONSE_HORS_DOEUVRES
//	@Router       /horsdoeuvres [post]
func CreateHorsDoeuvres(c *gin.Context) {
	// TODO: Add logic to reject unauthorized requests (and certainly do not deploy until all auth logic is wired up)
	response := V1_API_RESPONSE_HORS_DOEUVRES{}
	var status int
	var input models.HorsDoeuvres
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		status = http.StatusBadRequest
		response.Message = "\"option_name\" is required"
	} else {
		horsDoeuvres := []models.HorsDoeuvres{input}
		err := models.CreateHorsDoeuvres(&horsDoeuvres)
		if err != nil {
			status = http.StatusInternalServerError
			response.Message = "Internal server error"
			log.Println("Error inserting hors_doeuvres record: ", err.Error())
		} else {
			status = http.StatusCreated
			response.Message = "Created new hors doeuvres"
			response.Data.HorsDoeuvres = horsDoeuvres
		}
	}
	response.Status = status
	c.JSON(status, response)
}

// DeleteHorsDoeuvres deletes an hors doeuvres
//
//	@Summary      deletes an hors doeuvres
//	@Description  Deletes an hors doeuvres and returns a response to indicate success or failure
//	@Tags         hors doeuvres
//	@Accept       json
//	@Produce      json
//	@Success      200  {object}  V1_API_RESPONSE_HORS_DOEUVRES
//	@Failure      500  {object}  V1_API_RESPONSE_HORS_DOEUVRES
//	@Router       /horsdoeuvres [delete]
func DeleteHorsDoeuvres(c *gin.Context) {
	response := V1_API_DELETE_RESPONSE{}
	var status int
	id, _ := uuid.Parse(c.Param("id"))
	result, err := models.DeleteHorsDoeuvres(id)
	if err != nil {
		status = http.StatusInternalServerError
		response.Message = "Internal server error"
		log.Println("Error deleting hors doeuvres: ", err.Error())
	} else {
		status = http.StatusAccepted
		response.Message = "Deleted hors doeuvres"
		response.Data.DeletedRecords = int(*result)
	}
	response.Status = status
	c.JSON(status, response)
}
