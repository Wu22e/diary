package network

import (
	"rpc_server/config"
	"rpc_server/service"

	"github.com/gin-gonic/gin"
)

type Network struct {
	cfg *config.Config

	service *service.Service

	engine *gin.Engine
}

func NewNetwork(cfg *config.Config, service *service.Service) (*Network, error) {
	n := &Network{cfg: cfg, service: service, engine: gin.New()}

	return n, nil
}

func (n *Network) StartServer() {
	n.engine.Run(":8080")
}
