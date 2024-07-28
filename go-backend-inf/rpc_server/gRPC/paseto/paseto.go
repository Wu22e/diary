package paseto

import (
	"crypto/rand"
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

func (m *PasetoMaker) CreateNewToken(auth auth.AuthData) (string, error) {
	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)
	return m.Pt.Encrypt(m.Key, auth, randomBytes)
}

func (m *PasetoMaker) VerifyToken(token string) error {
	var auth auth.AuthData
	return m.Pt.Decrypt(token, m.Key, &auth, nil)
}
