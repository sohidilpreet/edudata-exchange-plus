package api

import (
	"app/config"
	"app/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SubmitApplication handles JSON-based submission
// @Summary Submit JSON Application
// @Tags Applications
// @Accept json
// @Produce json
// @Param data body models.Application true "Application Payload"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /applications [post]
func SubmitApplication(c *gin.Context) {
	var app models.Application
	if err := c.ShouldBindJSON(&app); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	query := `INSERT INTO applications (full_name, email, dob, program_applied) VALUES ($1, $2, $3, $4)`
	_, err := config.DB.Exec(query, app.FullName, app.Email, app.DOB, app.ProgramApplied)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Application submitted successfully"})
}
