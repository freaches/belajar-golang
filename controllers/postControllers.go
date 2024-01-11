package controllers

import (
	"errors"
	"micro-feature-pemilu/initializers"
	"micro-feature-pemilu/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostsCreate(c *gin.Context) {

	var body struct {
		Body   string
		Title  string
		TestID uint
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body, TestID: body.TestID}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Preload("Test").Find(&posts)

	c.JSON(200, gin.H{
		"post": posts,
	})
}

func PostShow(c *gin.Context) {

	id := c.Param("id")

	var post models.Post

	err := initializers.DB.First(&post, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(400, gin.H{
			"message": `not`,
		})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Body   string
		Title  string
		TestID uint
	}

	c.Bind(&body)

	var post models.Post
	result := initializers.DB.First(&post, id)

	if result == nil {
		c.Status(500)
		return
	}

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body, TestID: body.TestID})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.JSON(200, gin.H{
		"message": `succeded deleting a post`,
	})
}
