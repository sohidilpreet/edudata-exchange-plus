package api

import (
	"app/config"
	"app/internal/models"
	"app/internal/utils"
	"net/http"
	"os"

	"encoding/xml"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// SubmitApplicationXML handles XML-based PESC submission
// @Summary Submit XML Application (PESC-compliant only)
// @Tags Applications
// @Accept xml
// @Produce json
// @Param data body models.Application true "PESC XML Payload"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /applications/xml [post]
func SubmitApplicationXML(c *gin.Context) {
	var app models.Application
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read XML body"})
		return
	}

	// Write temp XML file for validation
	tempFile, err := ioutil.TempFile("", "pesc-*.xml")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Temp file error"})
		return
	}
	defer os.Remove(tempFile.Name())
	_ = ioutil.WriteFile(tempFile.Name(), body, 0644)

	// Validate against PESC schema
	xsdPath := "/app/internal/xmlschema/application.xsd"
	if err := utils.ValidateXMLWithXSD(tempFile.Name(), xsdPath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "❌ PESC XML validation failed"})
		return
	}

	// Parse XML into struct
	if err := xml.Unmarshal(body, &app); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "❌ Failed to parse XML"})
		return
	}

	// Store in DB
	query := `INSERT INTO applications (full_name, email, dob, program_applied) VALUES ($1, $2, $3, $4)`
	_, err = config.DB.Exec(query, app.FullName, app.Email, app.DOB, app.ProgramApplied)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "❌ Database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "✅ PESC XML Application submitted"})
}
