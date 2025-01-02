package repository

import (
	"backend/internal/domain"

	"github.com/jinzhu/gorm"
)

type TireRepository interface {
	GetAll() ([]domain.Tire, error)
	GetByID(id uint) (domain.Tire, error)
	Create(tire *domain.Tire) error
	Update(tire *domain.Tire) error
	Delete(id uint) error
}

type tireRepository struct {
	db *gorm.DB
}

func NewTireRepository(db *gorm.DB) TireRepository {
	return &tireRepository{db}
}

func (r *tireRepository) GetAll() ([]domain.Tire, error) {
	var tires []domain.Tire
	if err := r.db.Find(&tires).Error; err != nil {
		return nil, err
	}
	return tires, nil
}

func (r *tireRepository) GetByID(id uint) (domain.Tire, error) {
	var tire domain.Tire
	if err := r.db.First(&tire, id).Error; err != nil {
		return tire, err
	}
	return tire, nil
}

func (r *tireRepository) Create(tire *domain.Tire) error {
	return r.db.Create(tire).Error
}

func (r *tireRepository) Update(tire *domain.Tire) error {
	return r.db.Save(tire).Error
}

func (r *tireRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Tire{}, id).Error
}
