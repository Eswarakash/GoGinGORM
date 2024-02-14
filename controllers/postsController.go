package controllers

import (
	"github.com/Eswarakash/GoGinGORM/initializers"
	"github.com/Eswarakash/GoGinGORM/models"
	"github.com/gin-gonic/gin"
)

func GetPong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func PostsCreate(c *gin.Context) {
	//get data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	//create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//return it

}

func GetPosts(c *gin.Context) {
	//get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//response with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPost(c *gin.Context) {
	//get id from url
	id := c.Param("id")
	//get the post
	var post models.Post
	initializers.DB.First(&post, id)

	//response
	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	// get the id from url
	id := c.Param("id")

	//get data off body
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	//find the post where updateing
	var post models.Post
	initializers.DB.First(&post, id)

	//updste it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	//response
	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	//get the id from url
	id := c.Param("id")

	//delete the posts
	initializers.DB.Delete(&models.Post{}, id)

	//respone
	c.JSON(200, gin.H{
		"Deleted": id,
	})
}
