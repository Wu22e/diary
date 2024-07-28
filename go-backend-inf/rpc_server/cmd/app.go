package cmd

import (
	"rpc_server/config"
	"rpc_server/gRPC/client"
	"rpc_server/network"
	"rpc_server/repository"
	"rpc_server/service"
)

type App struct {
	cfg *config.Config

	gRPCClient *client.GRPCClient
	service    *service.Service
	repository *repository.Repository
	network    *network.Network
}

func NewApp(cfg *config.Config) {
	a := &App{cfg: cfg}

	var err error

	if a.repository, err = repository.NewRepository(cfg); err != nil {
		panic(err)
	} else if a.service, err = service.NewService(cfg, a.repository); err != nil {
		panic(err)
	} else if a.gRPCClient, err = client.NewGRPCClient(cfg); err != nil {
		panic(err)
	} else if a.network, err = network.NewNetwork(cfg, a.service, a.gRPCClient); err != nil {
		panic(err)
	} else {
		a.network.StartServer()
	}
}
