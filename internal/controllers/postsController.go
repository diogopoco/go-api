package controllers

import (
	dto "github.com/diogopoco/go-api/internal/dto/post"
	"github.com/diogopoco/go-api/internal/service"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {

	var body dto.CreatePostDto   

	c.Bind(&body)

	post, err := service.CreatePost(body)

	if err != nil {
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{
		"data": post,
	})
}

func GetPosts(c *gin.Context) {
	
	posts := service.GetAllPosts()

	c.JSON(200, gin.H{
		"data": posts,
	})
}

func GetPost(c *gin.Context) {

	id := c.Param("id")

	post := service.GetPost(id)

	c.JSON(200, gin.H{
		"data": post,
	})
}

func UpdatePost(c *gin.Context) {

	var body dto.UpdatePostDto

	c.Bind(&body)

	post := service.UpdatePost(body)

	c.JSON(200, gin.H{
		"data": post,
	})
}

func DeletePost(c *gin.Context) {

	id := c.Param("id")

	service.DeletePost(id)

	c.Status(200)
}
