package config

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
)

var (
	sqldb   *sql.DB
	redisDB *redis.Client
)

func configDataBase() {
	sqldb = configMysqlDatabase()
	redisDB = ConfigRedisDatabase()
}

func configMysqlDatabase() *sql.DB {
	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DBNAME")
	location := os.Getenv("MYSQL_LOCATION")

	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?loc=%s",
		username, password, host, port, dbName, url.QueryEscape(location))

	mysqlDB, err := sql.Open("mysql", datasource)
	if err != nil {
		fmt.Println("[CONFIG-DATABASE] Failed conntection to mysql: " + err.Error())
		os.Exit(1)

	}
	mysqlDB.SetMaxOpenConns(5)
	mysqlDB.SetMaxIdleConns(5)
	err = mysqlDB.Ping()
	if err != nil {
		fmt.Println("[CONFIG-DATABASE] Failed pinged to mysql: " + err.Error())
		os.Exit(1)
	}
	return mysqlDB
}

func ConfigRedisDatabase() *redis.Client {
	redisURI := os.Getenv("REDIS_URI")
	redisDB := os.Getenv("REDIS_DB")
	redisPASS := os.Getenv("REDIS_PASSWORD")
	db2, err := strconv.Atoi(redisDB)
	if err != nil {
		fmt.Println("[CONFIG-DATABASE] Failed Redis: " + err.Error())
		os.Exit(1)
	}
	client := redis.NewClient(&redis.Options{
		Addr:     redisURI,
		Password: redisPASS, // no password set
		DB:       db2,       // use default DB
	})
	status := client.Ping(context.Background())
	if status.Err() != nil {
		fmt.Println("[CONFIG-DATABASE] Failed Redis: " + err.Error())
		os.Exit(2)
	}
	return client
}
