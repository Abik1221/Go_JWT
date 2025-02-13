package main

import (
	"os"

	"github.com/NahomKeneni/go_jwt/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routes.UserRoutes(router)
	routes.AuthRoutes(router)

	router.Run(":" + port)
}
