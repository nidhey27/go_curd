package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nidhey27/go_curd_postgres/initializers"
	"github.com/nidhey27/go_curd_postgres/models"
)

func CreatePost(c *gin.Context) {

	var body models.Posts

	c.Bind(&body)

	post := models.Posts{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "",
			"error":   result.Error,
			"data":    make([]string, 0),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Post Created!",
		"error":   "",
		"data":    post,
	})
}

func GetPosts(c *gin.Context) {
	var posts []models.Posts
	initializers.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "All Posts",
		"error":   "",
		"data":    posts,
	})

}

func GetPost(c *gin.Context) {
	var posts []models.Posts

	id := c.Param("id")

	initializers.DB.Find(&posts, id)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Post Fetched",
		"error":   "",
		"data":    posts,
	})
}

func UpdatePost(c *gin.Context) {
	var body models.Posts
	id := c.Param("id")
	var posts []models.Posts
	c.Bind(&body)

	initializers.DB.Find(&posts, id)

	result := initializers.DB.Model(&posts).Updates(models.Posts{Title: body.Title, Body: body.Body})

	// post := models.Posts{Title: body.Title, Body: body.Body}

	// result := initializers.DB.Save(&post).Where("ID = ?", id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "",
			"error":   result.Error,
			"data":    make([]string, 0),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Post Updated!",
		"error":   "",
		"data":    posts,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	result := initializers.DB.Delete(&models.Posts{}, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "",
			"error":   result.Error,
			"data":    make([]string, 0),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Post Deleted!",
		"error":   "",
		"data":    make([]string, 0),
	})

}
