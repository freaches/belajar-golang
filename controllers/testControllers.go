package controllers

import (
	"errors"
	"micro-feature-pemilu/initializers"
	"micro-feature-pemilu/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TestsCreate(c *gin.Context) {

	var body struct {
		Name string
	}

	c.Bind(&body)

	post := models.Test{Name: body.Name}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func TestIndex(c *gin.Context) {
	var posts []models.Test
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"post": posts,
	})
}

func TestShow(c *gin.Context) {

	id := c.Param("id")

	var post models.Test

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

func TestsUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Name string
	}

	c.Bind(&body)

	var post models.Test
	result := initializers.DB.First(&post, id)

	if result == nil {
		c.Status(500)
		return
	}

	initializers.DB.Model(&post).Updates(models.Test{Name: body.Name})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeleteTest(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Test{}, id)

	c.JSON(200, gin.H{
		"message": `succeded deleting a post`,
	})
}
