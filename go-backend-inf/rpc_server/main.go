package main

import (
	"flag"
	"rpc_server/cmd"
	"rpc_server/config"
)

var configFlag = flag.String("config", "./config.toml", "config path")

func main() {
	flag.Parse()
	cfg := config.NewConfig(*configFlag)

	cmd.NewApp(cfg)
}
