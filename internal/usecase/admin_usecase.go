package usecase

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"fmt"
)

type AdminUsecase interface {
	GetAllAdmins() ([]domain.Admin, error)
	GetAdminByName(name, password string) (domain.Admin, error)
	CreateAdmin(admin *domain.Admin) error
	UpdateAdmin(admin *domain.Admin) error
	DeleteAdmin(id uint) error
}

type adminUsecase struct {
	adminRepo repository.AdminRepository
}

func NewAdminUsecase(adminRepo repository.AdminRepository) AdminUsecase {
	return &adminUsecase{adminRepo}
}

func (u *adminUsecase) GetAllAdmins() ([]domain.Admin, error) {
	return u.adminRepo.GetAll()
}

func (u *adminUsecase) GetAdminByName(name, password string) (domain.Admin, error) {
	admin, err := u.adminRepo.GetByName(name)
	if err != nil {
		return admin, err
	}

	if admin.Password != password {
		return admin, fmt.Errorf("password is not valid")
	}

	return admin, nil
}

func (u *adminUsecase) CreateAdmin(admin *domain.Admin) error {
	return u.adminRepo.Create(admin)
}

func (u *adminUsecase) UpdateAdmin(admin *domain.Admin) error {
	return u.adminRepo.Update(admin)
}

func (u *adminUsecase) DeleteAdmin(id uint) error {
	return u.adminRepo.Delete(id)
}
