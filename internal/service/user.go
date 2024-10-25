package service

import "future-path/internal/repository"

type IUserService interface {
}

type UserService struct {
	UserRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{userRepository}
}
