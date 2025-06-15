package main

import (
	"app/config"
	"app/internal/api"
	"app/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Application struct {
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	DOB            string `json:"dob"`
	ProgramApplied string `json:"program_applied"`
}

// SubmitApplication handles JSON-based application submission
// @Summary Submit JSON Application
// @Tags Applications
// @Accept json
// @Produce json
// @Param data body api.Application true "Application Payload"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /applications [post]
func SubmitApplication(c *gin.Context) {
	var app Application
	if err := c.ShouldBindJSON(&app); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	query := `
        INSERT INTO applications (full_name, email, dob, program_applied)
        VALUES ($1, $2, $3, $4);
    `

	_, err := config.DB.Exec(query, app.FullName, app.Email, app.DOB, app.ProgramApplied)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save application"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Application submitted successfully"})
}

func RegisterRoutes(r *gin.Engine) {
	r.POST("/login", api.Login)

	protected := r.Group("/")
	protected.Use(middleware.RequireAuth())
	{
		protected.POST("/applications", api.SubmitApplication)
		protected.POST("/applications/xml", api.SubmitApplicationXML)
	}
}
