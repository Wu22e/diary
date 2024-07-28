package network

import (
	"rpc_server/config"
	"rpc_server/gRPC/client"
	"rpc_server/service"

	"github.com/gin-gonic/gin"
)

type Network struct {
	cfg *config.Config

	service    *service.Service
	gRPCClient *client.GRPCClient

	engine *gin.Engine
}

func NewNetwork(cfg *config.Config, service *service.Service, gRPCClient *client.GRPCClient) (*Network, error) {
	n := &Network{cfg: cfg, service: service, engine: gin.New(), gRPCClient: gRPCClient}

	return n, nil
}

func (n *Network) StartServer() {
	n.engine.Run(":8080")
}
