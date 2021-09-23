package service

import "github.com/ValeryChapman/chat/pkg/repository"

type Service struct{}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
