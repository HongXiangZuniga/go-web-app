package config

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	sqldb *sql.DB
)

func configDataBase() {
	sqldb = configMysqlDatabase()
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
