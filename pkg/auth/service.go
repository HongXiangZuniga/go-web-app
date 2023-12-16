package auth

import (
	"crypto/sha512"
	"encoding/hex"

	"go.uber.org/zap"
)

type Service interface {
	Authorization(email, password string) (*bool, error)
}

type port struct {
	RepoSQL      Repository
	Logger       *zap.Logger
	Demo         bool
	DumyUser     string
	DumyPassword string
}

func NewService(repo Repository, Logger *zap.Logger, Demo bool, Dumyuser, DumyPassword string) Service {
	return &port{RepoSQL: repo, Logger: Logger, Demo: Demo, DumyUser: Dumyuser, DumyPassword: DumyPassword}
}

func (port *port) Authorization(email, password string) (*bool, error) {
	if port.Demo {
		result := true
		if port.DumyUser == email && port.DumyPassword == password {
			return &result, nil
		} else {
			result = false
			return &result, nil
		}
	}
	salt, err := port.RepoSQL.GetSalt(email)
	if err != nil {
		return nil, err
	}
	passwordEncrypted := password
	for i := 1; i <= 25000; i++ {
		passwordEncrypted = EncriptedSH512(passwordEncrypted + *salt)
	}
	isAuth, err := port.RepoSQL.AuthUser(email, hex.EncodeToString([]byte(passwordEncrypted)))
	if err != nil {
		return nil, err
	}
	return isAuth, nil
}

func EncriptedSH512(text string) string {
	sha512Hasher := sha512.New()
	sha512Hasher.Write([]byte(text))
	hashedPasswordBytes := sha512Hasher.Sum(nil)
	return (string(hashedPasswordBytes))
}
