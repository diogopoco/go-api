package interfaces

import (
	dto "github.com/diogopoco/go-api/internal/dto/post"
	"github.com/diogopoco/go-api/internal/entity"
)

type IPost interface {
	CreatePost(body dto.CreatePostDto) (entity.Post, error)
	GetAllPosts() []entity.Post
	GetPost(id string) entity.Post
	UpdatePost(body dto.UpdatePostDto) entity.Post
	DeletePost(id string)
}