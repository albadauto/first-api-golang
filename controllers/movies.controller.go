package controllers

import (
	"fmt"
	"strconv"

	"github.com/albadauto/projeto/database"
	"github.com/albadauto/projeto/models"
	"github.com/gin-gonic/gin"
)

func ShowAllMovies(c *gin.Context){
	db := database.GetDatabase()
	var movies []models.Movies
	err := db.Find(&movies).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "No movies" + err.Error(),
		})
	}

	c.JSON(200, movies)
}


func InsertMovie(c *gin.Context){
	db := database.GetDatabase()
	var movies models.Movies
	err := c.ShouldBindJSON(&movies)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "No movies" + err.Error(),
		})
	}
	
	err = db.Create(&movies).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "No movies" + err.Error(),
		})
	}

	fmt.Println(&movies)

}

func DeleteMovie(c *gin.Context){
	id := c.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Error: Id is just a int value " + err.Error(),
		})
		return

	}
	db := database.GetDatabase()
	var movies models.Movies
	err = db.First(&movies, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Not movie exist " + err.Error(),
		})
		return
	}

	c.JSON(200, movies)


}