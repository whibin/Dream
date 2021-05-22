// Package database 数据库连接
package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"Dream/conf"
)

var DB *gorm.DB

func init() {
	username := conf.Config.DB.Username
	password := conf.Config.DB.Password
	ip := conf.Config.DB.IP
	port := conf.Config.DB.Port
	database := conf.Config.DB.Database
	configuration := conf.Config.DB.Variables
	dbUrl := username + ":" + password + "@" + "(" + ip + ":" + port + ")/" + database + "?" + configuration
	var err error
	DB, err = gorm.Open(mysql.Open(dbUrl))
	if err != nil {
		panic("failed to connect database")
	}
	if DB.Error != nil {
		panic("database error")
	}
}
