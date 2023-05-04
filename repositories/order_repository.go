package repositories

import (
	"waysgallery/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	FindOrders() ([]models.Order, error)
	GetOrder(ID int) (models.Order, error)
	AddOrder(order models.Order) (models.Order, error)
	UpdateOrder(order models.Order) (models.Order, error)
	DeleteOrder(order models.Order, ID int) (models.Order, error)
}

func RepositoryOrder(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindOrders() ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Find(&orders).Error

	return orders, err
}

func (r *repository) GetOrder(ID int) (models.Order, error) {
	var order models.Order
	err := r.db.First(&order, ID).Error

	return order, err
}

func (r *repository) AddOrder(order models.Order) (models.Order, error) {
	err := r.db.Create(&order).Error

	return order, err
}

func (r *repository) UpdateOrder(order models.Order) (models.Order, error) {
	err := r.db.Save(&order).Error

	return order, err
}

func (r *repository) DeleteOrder(order models.Order, ID int) (models.Order, error) {
	err := r.db.Delete(&order).Error

	return order, err
}
