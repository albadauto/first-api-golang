package controllers

import (
	"github.com/albadauto/projeto/database"
	"github.com/albadauto/projeto/models"
	"github.com/gin-gonic/gin"
)

func ShowPeoples(c *gin.Context){
	db := database.GetDatabase()

	var peoples []models.Peoples
	err := db.Find(&peoples).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não existe pessoas cadastradas" + err.Error(),
		})

		return
	}

	c.JSON(200, peoples)
}


func InsertPeoples(c *gin.Context){
	db := database.GetDatabase()

	var peoples models.Peoples
	err := db.Create(&peoples).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Erroooo",
		})

		return
	}

	c.JSON(200, peoples)
}