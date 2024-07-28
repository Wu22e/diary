package server

import (
	"log"
	"net"
	"rpc_server/config"
	"rpc_server/gRPC/paseto"
	auth "rpc_server/gRPC/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	pasetoMaker    *paseto.PasetoMaker
	tokenVerifyMap map[string]auth.AuthData // implement memory db instead of repository db
}

func NewGRPCServer(cfg *config.Config) error {
	if lis, err := net.Listen("tcp", cfg.GRPC.URL); err != nil {
		return err
	} else {

		server := grpc.NewServer([]grpc.ServerOption{}...)

		auth.RegisterAuthServiceServer(server, &GRPCServer{
			pasetoMaker:    paseto.NewPasetoMaker(cfg),
			tokenVerifyMap: make(map[string]*auth.AuthData),
		})

		// grpc server register
		reflection.Register(server)

		go func() {
			log.Println("Start GRPC Server")
			if err = server.Serve(lis); err != nil {
				panic(err)
			}
		}()

		return nil
	}
}
