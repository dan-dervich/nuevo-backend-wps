package routes

import (
	"context"
	"fmt"

	"github.com/dan-dervich/nuevo-backend-wps/db"
	"github.com/gin-gonic/gin"
)

func NuevaCarta(c *gin.Context) {
	CartaModelo, err := db.ConectarBDCartas()
	if err != nil {
		panic(err)
	}
	var req db.Carta
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if len(req.CartaTitulo) < 1 {
		c.JSON(200, gin.H{"error": "title not long enough"})
		return
	}
	insertResult, err := CartaModelo.InsertOne(context.TODO(), req)
	if err != nil {
		panic(err)
	}
	fmt.Println(insertResult)
	c.JSON(200, gin.H{
		"_id": insertResult.InsertedID,
	})
}
