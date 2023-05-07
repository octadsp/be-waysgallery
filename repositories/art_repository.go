package repositories

import (
	"waysgallery/models"

	"gorm.io/gorm"
)

type ArtRepository interface {
	FindArts() ([]models.Art, error)
	GetArt(ID int) (models.Art, error)
	AddArt(art models.Art) (models.Art, error)
	UpdateArt(art models.Art) (models.Art, error)
	DeleteArt(art models.Art, ID int) (models.Art, error)
}

func RepositoryArt(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindArts() ([]models.Art, error) {
	var arts []models.Art
	err := r.db.Preload("User").Find(&arts).Error

	return arts, err
}

func (r *repository) GetArt(ID int) (models.Art, error) {
	var art models.Art
	err := r.db.Preload("User").First(&art, ID).Error

	return art, err
}

func (r *repository) AddArt(art models.Art) (models.Art, error) {
	err := r.db.Create(&art).Error

	return art, err
}

func (r *repository) UpdateArt(art models.Art) (models.Art, error) {
	err := r.db.Save(&art).Error

	return art, err
}

func (r *repository) DeleteArt(art models.Art, ID int) (models.Art, error) {
	err := r.db.Delete(&art).Error

	return art, err
}
