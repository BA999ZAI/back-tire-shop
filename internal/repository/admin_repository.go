package repository

import (
	"backend/internal/domain"

	"github.com/jinzhu/gorm"
)

type AdminRepository interface {
	GetAll() ([]domain.Admin, error)
	GetByName(name string) (domain.Admin, error)
	Create(admin *domain.Admin) error
	Update(admin *domain.Admin) error
	Delete(id uint) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db}
}

func (r *adminRepository) GetAll() ([]domain.Admin, error) {
	var admins []domain.Admin
	if err := r.db.Find(&admins).Error; err != nil {
		return nil, err
	}
	return admins, nil
}

func (r *adminRepository) GetByName(name string) (domain.Admin, error) {
	var admin domain.Admin
	if err := r.db.Where("username = ?", name).First(&admin).Error; err != nil {
		return admin, err
	}
	return admin, nil
}

func (r *adminRepository) Create(admin *domain.Admin) error {
	return r.db.Create(admin).Error
}

func (r *adminRepository) Update(admin *domain.Admin) error {
	return r.db.Save(admin).Error
}

func (r *adminRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Admin{}, id).Error
}
