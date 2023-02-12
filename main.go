package main

import (
	"github.com/gin-gonic/gin"
	"jwt-auth-service/controllers"
	"jwt-auth-service/database"
	"jwt-auth-service/middlewares"
)

func main() {
	// Initialize Database
	database.Connect("root:root@tcp(localhost:8889)/go_jwt_demo?parseTime=true")
	database.Migrate()

	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/login", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
