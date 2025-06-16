package api

import (
	"app/config"
	"app/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetApplications(c *gin.Context) {
	fmt.Println("🚨 Reached GetApplications handler!")
	rows, err := config.DB.Query(`SELECT id, full_name, email, dob, program_applied FROM applications`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "❌ Failed to query database"})
		return
	}
	defer rows.Close()

	var applications []models.ApplicationWithID

	for rows.Next() {
		var app models.ApplicationWithID
		if err := rows.Scan(&app.ID, &app.FullName, &app.Email, &app.DOB, &app.ProgramApplied); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "❌ Error scanning data"})
			return
		}
		applications = append(applications, app)
	}

	fmt.Println("Applications fetched from DB:", applications)
	c.JSON(http.StatusOK, applications)
}
