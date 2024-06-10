package controllers

import (
	"go-crud-example/initializers"
	"go-crud-example/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	var post models.Post

	initializers.DB.First(&post, c.Param("id"))

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	var oldPost models.Post
	initializers.DB.First(&oldPost, c.Param("id"))

	oldPost.Title = body.Title
	oldPost.Body = body.Body

	initializers.DB.Model(&oldPost).Updates(models.Post{Title: body.Title, Body: body.Body})

	c.JSON(200, gin.H{
		"post": body,
	})

}

func PostDelete(c *gin.Context) {

	var post models.Post
	initializers.DB.Delete(&post, c.Param("id"))

	c.JSON(200, gin.H{
		"message": "Post deleted",
	})
}
