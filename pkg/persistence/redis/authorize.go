package redis

import (
	"context"
	"time"

	"github.com/HongXiangZuniga/login-go/pkg/authorize"
	"github.com/go-redis/redis/v8"
)

type storage struct {
	redis *redis.Client
}

func NewAuthorizeRepo(db *redis.Client) authorize.RedisRepository {
	return &storage{db}
}

func (stg *storage) SetHash(hash, email string) error {
	status := stg.redis.Set(context.Background(), hash, email, time.Hour*5)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}
func (stg *storage) Authorize(hash string) (*string, error) {
	var email string
	status := stg.redis.Get(context.Background(), hash)
	if status.Err() != nil {
		return nil, status.Err()
	}
	err := status.Scan(&email)
	if err != nil {
		return nil, err
	}
	return &email, nil
}
