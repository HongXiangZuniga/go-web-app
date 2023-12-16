package config

import (
	"github.com/HongXiangZuniga/login-go/pkg/auth"
	"github.com/HongXiangZuniga/login-go/pkg/persistence/mysql"
)

var (
	AuthSQLRepo auth.Repository
)

func configRepository() {
	AuthSQLRepo = configAuthSQLRepo()
}

func configAuthSQLRepo() auth.Repository {
	return mysql.NewAuthRepository(sqldb)
}
