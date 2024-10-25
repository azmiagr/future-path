package service

import "future-path/internal/repository"

type Service struct {
	UserService   IUserService
	BeritaService IBeritaService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		UserService:   NewUserService(repository.UserRepository),
		BeritaService: NewBeritaService(repository.BeritaRepository),
	}
}
