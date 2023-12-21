package authorize

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"go.uber.org/zap"
)

type Service interface {
	SetHash(email string) (*string, error)
	Authorize(hash string) (*bool, error)
}

type port struct {
	repo   RedisRepository
	logger *zap.Logger
}

func NewService(repo RedisRepository, logger *zap.Logger) Service {
	return &port{repo, logger}
}

func (port *port) SetHash(email string) (*string, error) {
	hash, err := generateRandomHash()
	if err != nil {
		return nil, err
	}
	err = port.repo.SetHash(*hash, email)
	if err != nil {
		return nil, err
	}
	fmt.Println(*hash)
	return hash, nil
}

func (port *port) Authorize(hash string) (*bool, error) {
	authorize := true
	email, err := port.repo.Authorize(hash)
	if err != nil || email == nil {
		authorize = false
		if err.Error() == "redis: nil" {
			return &authorize, nil
		}
		return &authorize, err
	}
	return &authorize, nil
}

func generateRandomHash() (*string, error) {
	randomBytes := make([]byte, 32) // Tamaño arbitrario para este ejemplo (puedes ajustarlo según tus necesidades)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}
	hasher := sha256.New()
	_, err = hasher.Write(randomBytes)
	if err != nil {
		return nil, err
	}
	hashInBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashInBytes)
	return &hashString, nil
}
