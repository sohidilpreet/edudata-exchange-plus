package main

import (
	"app/config"
	_ "app/docs"

	"log"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title EduData Exchange+ API
// @version 1.0
// @description Process student applications using secure JSON/XML APIs with PESC validation.
// @contact.name Dilpreet Singh Sohi
// @contact.email sohidilpreet1999@gmail.com
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
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
	RegisterRoutes(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
