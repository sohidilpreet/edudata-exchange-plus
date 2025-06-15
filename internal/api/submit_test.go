package api

import (
	"app/config"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func TestSubmitApplication(t *testing.T) {
	// Load local test environment
	err := godotenv.Load("../../.env.test")
	if err != nil {
		t.Fatalf("❌ Failed to load .env.test: %v", err)
	}

	// Init DB
	if err := config.InitDB(); err != nil {
		t.Fatalf("❌ Failed to init DB: %v", err)
	}

	config.CreateApplicationTable()

	// Setup router
	router := gin.Default()
	router.POST("/applications", SubmitApplication)

	body := `{"full_name":"Test User","email":"test@example.com","dob":"1990-01-01","program_applied":"CS"}`
	req, _ := http.NewRequest("POST", "/applications", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("❌ Expected 200 OK, got %d", resp.Code)
	}
}
