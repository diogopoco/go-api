package main

import (
	"github.com/diogopoco/go-api/internal/entity"
	"github.com/diogopoco/go-api/internal/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&entity.Post{})
}