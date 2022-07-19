package controllers

import (
	"strconv"

	"github.com/albadauto/projeto/database"
	"github.com/albadauto/projeto/models"
	"github.com/gin-gonic/gin"
)

func ShowBook(c *gin.Context){
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID Has to be a integer",
		})
		return 
	}

	db := database.GetDatabase()

	var book models.Books
	err = db.First(&book, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find book" + err.Error(),
		})
		return
	}

	c.JSON(200, book)

}

func CreateBook(c *gin.Context){
	db := database.GetDatabase()

	var book models.Books

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	c.JSON(200, book)

}

func ShowBooks(c *gin.Context){
	db := database.GetDatabase()
	var books []models.Books
	err := db.Find(&books).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot show BOOKS: " + err.Error(),
		})
		return
	}

	c.JSON(200, books)
}