package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

type CrearCartaReq struct {
	CartaTitulo       string `json:"carta_titulo"`
	CartaDesc         string `json:"carta_desc"`
	CartaImagen       string `json:"carta_imagen"`
	MasInfoCarta      string `json:"mas_info_carta"`
	CategoriaCarta    string `json:"categoria_carta"`
	SubCategoriaCarta string `json:"sub_categoria_carta"`
}

func NuevaCarta(c *gin.Context) {
	CartaModelo, err := ConectarBDCartas()
	if err != nil {
		panic(err)
	}
	var req CrearCartaReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	crearCarta := Carta{
		CartaTitulo:       req.CartaTitulo,
		CartaDesc:         req.CartaDesc,
		CartaImagen:       req.CartaImagen,
		MasInfoCarta:      req.MasInfoCarta,
		CategoriaCarta:    req.CategoriaCarta,
		SubCategoriaCarta: req.SubCategoriaCarta,
	}
	insertResult, err := CartaModelo.InsertOne(context.TODO(), crearCarta)
	if err != nil {
		panic(err)
	}
	fmt.Println(insertResult)
	c.JSON(200, gin.H{
		"_id": insertResult.InsertedID,
	})
}
