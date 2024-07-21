package network

import (
	"crud_server/service"

	"github.com/gin-gonic/gin"
)

type Network struct {
	engine *gin.Engine

	service *service.Service
}

// Among the network frameworks mux, echo, and gin, we will use gin, which is the most widely used.
func NewNetwork(service *service.Service) *Network {
	r := &Network{
		engine: gin.New(),
	}

	newUserRouter(r, service.User)

	return r
}

func (n *Network) ServerStart(port string) error {
	return n.engine.Run(port)
}
