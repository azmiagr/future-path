package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository   IUserRepository
	BeritaRepository IBeritaRepository
	AdminRepository  IAdminRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository:   NewUserRepository(db),
		BeritaRepository: NewBeritaRepository(db),
		AdminRepository:  NewAdminRepository(db),
	}
}
