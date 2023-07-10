package main

import (
	"github.com/diogopoco/go-api/internal/controllers"
	"github.com/diogopoco/go-api/internal/initializers"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/Post", controllers.CreatePost)
	r.GET("/Post", controllers.GetPosts)
	r.GET("/Post/:id", controllers.GetPost)
	r.PUT("/Post", controllers.UpdatePost)
	r.DELETE("/Post/:id", controllers.DeletePost)

	r.Run() // listen and serve on 0.0.0.0:8080
}