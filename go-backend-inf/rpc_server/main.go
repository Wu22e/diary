package main

import (
	"flag"
	"rpc_server/config"
)

var configFlag = flag.String("config", "./config.toml", "config path")

func main() {
	flag.Parse()
	config.NewConfig(*configFlag)
}
