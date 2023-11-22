package main

import (
	"example.com/m/controllers"
	"example.com/m/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVar()
	initializers.ConnectToDb()
	initializers.SyncDB()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", controllers.Validate)
	r.Run()
}
