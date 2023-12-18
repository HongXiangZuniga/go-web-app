package config

import (
	"github.com/HongXiangZuniga/login-go/pkg/auth"
	"github.com/HongXiangZuniga/login-go/pkg/authorize"
	"github.com/HongXiangZuniga/login-go/pkg/persistence/mysql"
	"github.com/HongXiangZuniga/login-go/pkg/persistence/redis"
)

var (
	AuthSQLRepo              auth.Repository
	AuthorizeRedisRepository authorize.RedisRepository
)

func configRepository() {
	AuthSQLRepo = configAuthSQLRepo()
	AuthorizeRedisRepository = configAuthorizeRedisRepository()
}

func configAuthSQLRepo() auth.Repository {
	return mysql.NewAuthRepository(sqldb)
}

func configAuthorizeRedisRepository() authorize.RedisRepository {
	return redis.NewAuthorizeRepo(redisDB)
}
