package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/dan-dervich/nuevo-backend-wps/routes"
)

func main() {
	router := gin.New()
	router.SetTrustedProxies([]string{"localhost:3000", "localhost:8080"})
	router.Use(cors.Default())
	router1 := router.Group("/carta")
	{
		go router1.POST("/nueva", routes.NuevaCarta)
	}
	router.Run(":8080")
}
