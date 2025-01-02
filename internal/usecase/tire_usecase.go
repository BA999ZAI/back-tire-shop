package usecase

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type TireUsecase interface {
	GetAllTires() ([]domain.Tire, error)
	GetTireByID(id uint) (domain.Tire, error)
	CreateTire(tire *domain.Tire) error
	UpdateTire(tire *domain.Tire) error
	DeleteTire(id uint) error
}

type tireUsecase struct {
	tireRepo repository.TireRepository
}

func NewTireUsecase(tireRepo repository.TireRepository) TireUsecase {
	return &tireUsecase{tireRepo}
}

func (u *tireUsecase) GetAllTires() ([]domain.Tire, error) {
	return u.tireRepo.GetAll()
}

func (u *tireUsecase) GetTireByID(id uint) (domain.Tire, error) {
	return u.tireRepo.GetByID(id)
}

func (u *tireUsecase) CreateTire(tire *domain.Tire) error {
	return u.tireRepo.Create(tire)
}

func (u *tireUsecase) UpdateTire(tire *domain.Tire) error {
	return u.tireRepo.Update(tire)
}

func (u *tireUsecase) DeleteTire(id uint) error {
	return u.tireRepo.Delete(id)
}
