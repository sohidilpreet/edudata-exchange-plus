// @title EduData Exchange+ API
// @version 1.0
// @description API for processing student applications via JSON and XML
// @termsOfService http://swagger.io/terms/

// @contact.name Dilpreet Singh Sohi
// @contact.email sohidilpreet1999@gmail.com

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"app/config"
	_ "app/docs"
	"log"
	"time"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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

	RegisterRoutes(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
