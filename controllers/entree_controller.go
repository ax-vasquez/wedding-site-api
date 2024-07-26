package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ax-vasquez/wedding-site-api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Get a list of entrees
//
// If a user ID is specified, then this will return the list of Entrees containing
// one item (data for the entree they have selected), or zero items if no selection
// has been made, yet. If no user ID is specified, then data for all possible Entrees
// is returned.
func GetEntrees(c *gin.Context) {
	idStr := c.Param("id")
	var response V1_API_RESPONSE
	var status int
	var hors_doeuvres []models.Entree
	if len(idStr) > 0 {
		id, err := uuid.Parse(idStr)
		if err != nil {
			status = http.StatusInternalServerError
		} else {
			status = http.StatusOK
			hors_doeuvres, err = models.FindEntreesForUser(id)
			if err != nil {
				status = http.StatusInternalServerError
			}
		}
	} else {
		hors_doeuvres = models.FindEntrees()
		status = http.StatusOK
	}
	response.Status = status
	response.Data = gin.H{
		"entrees": hors_doeuvres}
	c.JSON(status, response)
}

// Controller to handle creating a new entree
//
// The route that uses this controller must be protected so that
// only site admins can use this endpoint. All other requests
// should be rejected.
func CreateEntree(c *gin.Context) {
	// TODO: Add logic to reject unauthorized requests (and certainly do not deploy until all auth logic is wired up)
	response := V1_API_RESPONSE{}
	var status int
	var input models.Entree
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		status = http.StatusBadRequest
		response.Message = "\"option_name\" is required"
	} else {
		result, err := models.CreateEntrees(&[]models.Entree{input})
		if err != nil {
			status = http.StatusInternalServerError
			response.Message = "Internal server error"
			log.Println("Error inserting entrees record: ", err.Error())
		} else {
			status = http.StatusCreated
			response.Message = "Created entree"
			response.Data = gin.H{"records": result}
		}
	}
	fmt.Println("STATUS IN CONTROLLER: ", status)
	response.Status = status
	c.JSON(status, response)
}

// Delete an entree
func DeleteEntree(c *gin.Context) {
	// TODO: Add logic to reject unauthorized requests (and certainly do not deploy until all auth logic is wired up)
	response := V1_API_RESPONSE{}
	var status int
	id, _ := uuid.Parse(c.Param("id"))
	result, err := models.DeleteEntree(id)
	if err != nil {
		status = http.StatusInternalServerError
		response.Message = "Internal server error"
		log.Println("Error deleting hors doeuvres: ", err.Error())
	} else {
		status = http.StatusAccepted
		response.Message = "Deleted hors doeuvres"
		response.Data = gin.H{"records": result}
	}
	response.Status = status
	c.JSON(status, response)
}
