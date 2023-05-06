package repositories

import (
	"waysgallery/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	FindPosts() ([]models.PostResponse, error)
	GetPost(ID int) (models.Post, error)
	AddPost(post models.Post) (models.Post, error)
	UpdatePost(post models.Post) (models.Post, error)
	DeletePost(post models.Post, ID int) (models.Post, error)
}

func RepositoryPost(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindPosts() ([]models.PostResponse, error) {
	var posts []models.PostResponse
	err := r.db.Preload("Photos").Preload("User").Find(&posts).Error

	return posts, err
}

func (r *repository) GetPost(ID int) (models.Post, error) {
	var post models.Post
	err := r.db.Preload("Photos").Preload("User").First(&post, ID).Error

	return post, err
}

func (r *repository) AddPost(post models.Post) (models.Post, error) {
	err := r.db.Create(&post).Error

	return post, err
}

func (r *repository) UpdatePost(post models.Post) (models.Post, error) {
	err := r.db.Save(&post).Error

	return post, err
}

func (r *repository) DeletePost(post models.Post, ID int) (models.Post, error) {
	err := r.db.Delete(&post).Error

	return post, err
}
