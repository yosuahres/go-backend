package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yosuahres/go-backend/controllers"
	"github.com/yosuahres/go-backend/initializers"
	"github.com/yosuahres/go-backend/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main()	{
	r := gin.Default()
	
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	
	r.Run() // listen and serve on 0.0.0.0:8080
}