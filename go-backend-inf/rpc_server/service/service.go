package service

import (
	"rpc_server/config"
	"rpc_server/repository"
)

type Service struct {
	cfg *config.Config

	repository *repository.Repository
}

func NewService(cfg *config.Config, repository *repository.Repository) (*Service, error) {
	s := &Service{cfg: cfg, repository: repository}

	return s, nil
}

func (s *Service) CreateAuth(name string) (interface{}, error) {
	return s.repository.CreateAuth(name)
}
