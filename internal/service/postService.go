package service

import (
	dto "github.com/diogopoco/go-api/internal/dto/post"
	"github.com/diogopoco/go-api/internal/entity"
	"github.com/diogopoco/go-api/internal/initializers"
)

func CreatePost(body dto.CreatePostDto) (entity.Post, error) {
	post := entity.Post{Title: body.Title, Descriptiom: body.Description, CreatedBy: body.CreatedBy }
	result := initializers.DB.Create(&post)

	return post, result.Error
}

func GetAllPosts() []entity.Post{
	var posts []entity.Post
	initializers.DB.Find(&posts)
	return posts
}

func GetPost(id string) entity.Post{
	var post entity.Post
	initializers.DB.First(&post, id)
	return post
}

func UpdatePost(body dto.UpdatePostDto) entity.Post {
	var post entity.Post
	initializers.DB.First(&post, body.Id)

	initializers.DB.Model(&post).Updates(entity.Post{Title:body.Title, Descriptiom: body.Description})
	return post
}

func DeletePost(id string) {
	initializers.DB.Delete(&entity.Post{}, id)
}