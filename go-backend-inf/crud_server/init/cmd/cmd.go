package cmd

import (
	"crud_server/config"
	"crud_server/network"
	"crud_server/repository"
	"crud_server/service"
)

type Cmd struct {
	config *config.Config

	network    *network.Network
	repository *repository.Repository
	service    *service.Service
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filePath),
	}

	c.repository = repository.NewRepository()
	c.service = service.NewService(c.repository)
	c.network = network.NewNetwork(c.service)

	c.network.ServerStart(c.config.Server.Port)

	return c
}
