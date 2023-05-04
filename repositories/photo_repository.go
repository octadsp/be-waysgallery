package repositories

import (
	"waysgallery/models"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	FindPhotos() ([]models.Photo, error)
	GetPhoto(ID int) (models.Photo, error)
	AddPhoto(photo models.Photo) (models.Photo, error)
	UpdatePhoto(photo models.Photo) (models.Photo, error)
	DeletePhoto(photo models.Photo, ID int) (models.Photo, error)
}

func RepositoryPhoto(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindPhotos() ([]models.Photo, error) {
	var photos []models.Photo
	err := r.db.Find(&photos).Error

	return photos, err
}

func (r *repository) GetPhoto(ID int) (models.Photo, error) {
	var photo models.Photo
	err := r.db.First(&photo, ID).Error

	return photo, err
}

func (r *repository) AddPhoto(photo models.Photo) (models.Photo, error) {
	err := r.db.Create(&photo).Error

	return photo, err
}

func (r *repository) UpdatePhoto(photo models.Photo) (models.Photo, error) {
	err := r.db.Save(&photo).Error

	return photo, err
}

func (r *repository) DeletePhoto(photo models.Photo, ID int) (models.Photo, error) {
	err := r.db.Delete(&photo).Error

	return photo, err
}
