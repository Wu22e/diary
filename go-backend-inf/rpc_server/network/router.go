package network

import (
	"rpc_server/config"
	"rpc_server/service"
)

type Network struct {
	cfg *config.Config

	service *service.Service
}

func NewNetwork(cfg *config.Config, service *service.Service) (*Network, error) {
	n := &Network{cfg: cfg, service: service}

	return n, nil
}
