package paseto

import (
	"rpc_server/config"
	auth "rpc_server/gRPC/proto"

	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	Pt  *paseto.V2
	Key []byte
}

func NewPasetoMaker(cfg *config.Config) *PasetoMaker {
	return &PasetoMaker{
		Pt:  paseto.NewV2(),
		Key: []byte(cfg.Paseto.Key),
	}
}

func (m *PasetoMaker) CreateNewToken(auth *auth.AuthData) (string, error) {
	return "", nil
}

func (m *PasetoMaker) VerifyToken(token string) error {
	return nil
}
