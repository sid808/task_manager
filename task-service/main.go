// @title Task Service API
// @version 1.0
// @description API documentation for the Task Service
// @host localhost:8080
// @BasePath /

// @schemes http

package main

import (
	"log"
	_ "task-service/docs" // Import your generated Swagger docs
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	"task-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.RegisterTaskRoutes(r)

	log.Println("Task Service running on port 8081")
	r.Run(":8081") // Runs on port 8081
}
