package controllers

import (
	"log"
	"net/http"

	"github.com/ax-vasquez/wedding-site-api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EntreeData struct {
	Entrees []models.Entree `json:"entrees"`
}

type V1_API_RESPONSE_ENTREE struct {
	V1_API_RESPONSE
	Data EntreeData `json:"data"`
}

// Get a list of entrees
//
// If a user ID is specified, then this will return the list of Entrees containing
// one item (data for the entree they have selected), or zero items if no selection
// has been made, yet. If no user ID is specified, then data for all possible Entrees
// is returned.
func GetEntrees(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	response := V1_API_RESPONSE_ENTREE{}
	var status int
	var entrees []models.Entree
	// If no error occurred, the parse was successful, meaning a UUID was found and results will be filtered for the given user
	if err == nil {
		status = http.StatusOK
		entrees, err = models.FindEntreesForUser(id)
		if err != nil {
			status = http.StatusInternalServerError
		}
		// If an error occurred, we ignore it and assume it's because there was no ID in the path - all results will be returned
	} else {
		entrees, err = models.FindEntrees()
		if err != nil {
			status = http.StatusInternalServerError
		} else {
			status = http.StatusOK
		}
	}
	response.Status = status
	response.Data.Entrees = entrees
	c.JSON(status, response)
}

// Controller to handle creating a new entree
//
// The route that uses this controller must be protected so that
// only site admins can use this endpoint. All other requests
// should be rejected.
func CreateEntree(c *gin.Context) {
	// TODO: Add logic to reject unauthorized requests (and certainly do not deploy until all auth logic is wired up)
	response := V1_API_RESPONSE_ENTREE{}
	var status int
	var input models.Entree
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		status = http.StatusBadRequest
		response.Message = "\"option_name\" is required"
	} else {
		entrees := []models.Entree{input}
		err := models.CreateEntrees(&entrees)
		if err != nil {
			status = http.StatusInternalServerError
			response.Message = "Internal server error"
			log.Println("Error inserting entrees record: ", err.Error())
		} else {
			status = http.StatusCreated
			response.Message = "Created entree"
			response.Data.Entrees = entrees
		}
	}
	response.Status = status
	c.JSON(status, response)
}

// Delete an entree
func DeleteEntree(c *gin.Context) {
	// TODO: Add logic to reject unauthorized requests (and certainly do not deploy until all auth logic is wired up)
	response := V1_API_DELETE_RESPONSE{}
	var status int
	id, _ := uuid.Parse(c.Param("id"))
	result, err := models.DeleteEntree(id)
	if err != nil {
		status = http.StatusInternalServerError
		response.Message = "Internal server error"
	} else {
		status = http.StatusAccepted
		response.Message = "Deleted entree"
		response.Data = DeleteRecordResponse{
			DeletedRecords: int(*result),
		}
	}
	response.Status = status
	c.JSON(status, response)
}
