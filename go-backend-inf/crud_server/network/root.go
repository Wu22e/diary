package network

import "github.com/gin-gonic/gin"

type Network struct {
	engine *gin.Engine
}

// Among the network frameworks mux, echo, and gin, we will use gin, which is the most widely used.
func NewNetwork() *Network {
	r := &Network{
		engine: gin.New(),
	}

	newUserRouter(r)

	return r
}

func (n *Network) ServerStart(port string) error {
	return n.engine.Run(port)
}
