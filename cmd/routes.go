package main

import (
	"app/internal/api"
	"app/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/login", api.Login)
	r.POST("/applications/validate", api.ValidatePESCXML)

	protected := r.Group("/")
	protected.Use(middleware.RequireAuth())
	{
		protected.POST("/applications", api.SubmitApplication)
		protected.POST("/applications/xml", api.SubmitApplicationXML)
		protected.GET("/applications", api.GetApplications)
	}
}
