package model

import (
	"fmt"
	"scaffolding/conf"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database(conf *conf.Configure) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Mysql.Username, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(conf.Mysql.MaxIdleConn)
	sqlDB.SetMaxOpenConns(conf.Mysql.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Duration(conf.Mysql.ConnMaxLifetime))
	DB = db
	migration()
}
