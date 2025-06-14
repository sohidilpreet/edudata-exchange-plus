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
