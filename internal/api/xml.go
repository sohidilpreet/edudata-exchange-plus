package api

import (
	"encoding/xml"
	"net/http"

	"app/config"

	"github.com/gin-gonic/gin"
)

type XMLApplication struct {
	XMLName        xml.Name `xml:"Application"`
	FullName       string   `xml:"FullName"`
	Email          string   `xml:"Email"`
	DOB            string   `xml:"DOB"`
	ProgramApplied string   `xml:"ProgramApplied"`
}

// SubmitApplicationXML handles XML-based submission
// @Summary Submit XML Application
// @Tags Applications
// @Accept xml
// @Produce json
// @Param data body models.Application true "XML Payload"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /applications/xml [post]
func SubmitApplicationXML(c *gin.Context) {
	var app XMLApplication

	if err := c.ShouldBindXML(&app); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid XML"})
		return
	}

	query := `INSERT INTO applications (full_name, email, dob, program_applied)
	          VALUES ($1, $2, $3, $4)`

	_, err := config.DB.Exec(query, app.FullName, app.Email, app.DOB, app.ProgramApplied)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Application (XML) submitted successfully"})
}
