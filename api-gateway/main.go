package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http/httputil"
	"net/url"
	"log"
)

func main() {
	r := gin.Default()

	// Allow CORS for Swagger UI
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Proxy routes
	taskService := createReverseProxy("http://localhost:8081")
	r.Any("/tasks", taskService)
	r.Any("/tasks/:id", taskService)

	log.Println("API Gateway running on port 8080")
	r.Run(":8080")
}

// Reverse proxy handler
func createReverseProxy(target string) gin.HandlerFunc {
	url, err := url.Parse(target)
	if err != nil {
		log.Fatalf("Invalid proxy target: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	return func(c *gin.Context) {
		c.Request.URL.Scheme = url.Scheme
		c.Request.URL.Host = url.Host
		c.Request.Host = url.Host
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
