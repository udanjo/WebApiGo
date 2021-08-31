package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ueverson/WebApiGo/database"
	"github.com/ueverson/WebApiGo/models"
)

func ShowMovie(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	var movie models.Movie
	err = db.First(&movie, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find movie: " + err.Error(),
		})
		return
	}

	c.JSON(200, movie)
}

func CreateMovie(c *gin.Context) {
	db := database.GetDatabase()

	var movie models.Movie

	err := c.ShouldBindJSON(&movie)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind json movie: " + err.Error(),
		})
		return
	}

	err = db.Create(&movie).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create movie: " + err.Error(),
		})
		return
	}

	c.JSON(201, movie)
}

func ShowMovies(c *gin.Context) {
	db := database.GetDatabase()

	var movies []models.Movie

	err := db.Find(&movies).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list movie: " + err.Error(),
		})
		return
	}

	c.JSON(200, movies)
}

func UpdateMovie(c *gin.Context) {
	db := database.GetDatabase()

	var movie models.Movie

	err := c.ShouldBindJSON(&movie)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind json movie: " + err.Error(),
		})
		return
	}

	err = db.Save(&movie).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update movie: " + err.Error(),
		})
		return
	}

	c.JSON(201, movie)
}

func DeleteMovie(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Movie{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete movie: " + err.Error(),
		})
		return
	}

	c.Status(204)
}
