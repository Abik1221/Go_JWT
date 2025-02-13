package routes

import "github.com/gin-gonic/gin"

func UserRoutes(incoming_route *gin.Engine) {
	incoming_route.Use(middleware.Authenticate())
	incoming_route.GET("/users", controllers.GetUsers())
	incoming_route.GET("/user/:id", controllers.GetUser())
}
