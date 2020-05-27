package model

import (
	"database/sql"
	"fmt"
	"self-site/setting"
)

var db *sql.DB

func Setup() {
	var err error
	driverName, dataSourceName := loadDBConfig()
	fmt.Println(driverName, dataSourceName, err)
	// db, err = sql.Open(driverName, dataSourceName)

	// if err = db.Ping(); err != nil {
	// 	log.Panic(err)
	// }
}

func CloseDB() {
	defer db.Close()
}

func loadDBConfig() (string, string) {
	connection := setting.Config.DB.Host
	host := setting.Config.DB.Host
	port := setting.Config.DB.Port
	database := setting.Config.DB.Database
	username := setting.Config.DB.Username
	password := setting.Config.DB.Password
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		username,
		password,
		host,
		port,
		database)
	return connection, dataSourceName
}
