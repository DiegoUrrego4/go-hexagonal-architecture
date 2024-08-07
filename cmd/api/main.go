package main

import (
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"time-tracker/cmd/server"
)

func main() {
	addCfg := ":8080"
	mysqlCfg := mysql.Config{
		User:      "root",
		Passwd:    "rootpassword",
		Net:       "tcp",
		Addr:      "localhost:3306",
		DBName:    "userdb",
		ParseTime: true,
	}

	cfg := server.ServerChi{Addr: addCfg, MySQLDSN: mysqlCfg.FormatDSN()}
	srv := server.New(cfg)

	srv.Run()
}
