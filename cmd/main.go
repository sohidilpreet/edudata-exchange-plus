package main

import (
	"app/config"
	"app/internal/api"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	var err error
	for i := 1; i <= 10; i++ {
		err = config.InitDB()
		if err == nil {
			break
		}
		log.Printf("❌ DB connection failed (try %d): %v", i, err)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatalf("❌ DB connection failed after retries: %v", err)
	}

	log.Println("✅ Connected to DB")
	config.CreateApplicationTable()
	log.Println("✅ applications table is ready")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/applications", api.SubmitApplication)
	r.POST("/applications/xml", api.SubmitApplicationXML) // 🔥 THE CRITICAL LINE

	r.Run(":8080")
}
