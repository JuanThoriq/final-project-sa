package routes

import (
	"final-project-sa-be/controllers"

	"final-project-sa-be/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Public routes
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	// Protected routes
	auth := router.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.POST("/cv", controllers.CreateCV)
		auth.GET("/cv", controllers.GetCVs)
		auth.GET("/cv/:id", controllers.GetCVByID)
		auth.PUT("/cv/:id", controllers.UpdateCV)
		auth.DELETE("/cv/:id", controllers.DeleteCV)
		auth.POST("/cv/:id/skills", controllers.AddSkillsToCV)

	}

	return router
}
