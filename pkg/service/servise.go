package service

import service "github.com/vintkor/todo/pkg/repository"

type Auth struct {
}

type User struct {
}

type Service struct {
	Auth
	User
}

func NewService(repos *service.Repository) *Service {
	return &Service{}
}
