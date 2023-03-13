package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iamsushank/jwt-go/controllers"
	"github.com/iamsushank/jwt-go/middlewares"
	"github.com/iamsushank/jwt-go/models"
)

func main() {
	models.ConnectDatabase()
	r := gin.Default()
	public := r.Group("/api")
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleWare())
	protected.GET("/user", controllers.CurrentUser)

	r.Run(":8080")

}
