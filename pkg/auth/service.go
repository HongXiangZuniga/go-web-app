package auth

import (
	"go.uber.org/zap"
)

type Service interface {
	Authorization(user, password string) (bool, error)
}

type port struct {
	Logger       *zap.Logger
	Demo         bool
	DumyUser     string
	DumyPassword string
}

func NewService(Logger *zap.Logger, Demo bool, Dumyuser, DumyPassword string) Service {
	return &port{Logger: Logger, Demo: Demo, DumyUser: Dumyuser, DumyPassword: DumyPassword}
}

func (port *port) Authorization(user, password string) (bool, error) {
	if port.Demo {
		if port.DumyUser == user && port.DumyPassword == password {
			return true, nil
		} else {
			return false, nil
		}
	}
	return false, nil
}
