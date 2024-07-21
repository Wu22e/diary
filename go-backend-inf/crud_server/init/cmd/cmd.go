package cmd

import (
	"crud_server/config"
	"crud_server/network"
	"fmt"
)

type Cmd struct {
	config  *config.Config
	network *network.Network
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config:  config.NewConfig(filePath),
		network: network.NewNetwork(),
	}

	c.network.ServerStart(c.config.Server.Port)

	// not reached until channel(gin network) close
	fmt.Println(c.config.Server.Port)

	return c
}
