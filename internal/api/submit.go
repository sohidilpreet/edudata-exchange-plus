package api

import (
	"net/http"

	"app/config"

	"github.com/gin-gonic/gin"
)

type Application struct {
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	DOB            string `json:"dob"`
	ProgramApplied string `json:"program_applied"`
}

func SubmitApplication(c *gin.Context) {
	var app Application

	if err := c.ShouldBindJSON(&app); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	query := `INSERT INTO applications (full_name, email, dob, program_applied)
	          VALUES ($1, $2, $3, $4)`

	_, err := config.DB.Exec(query, app.FullName, app.Email, app.DOB, app.ProgramApplied)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Application submitted successfully"})
}
