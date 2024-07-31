package core

import (
	"github.com/Linxhhh/easy-doc/global"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitMysql() *gorm.DB {

	if global.Config.Mysql.Host == "" {
		logrus.Warnln("未配置 Mysql 主机名，取消 Mysql 连接！")
		return nil
	}
	
	// 设置 Gorm 日志
	logLevel := logger.Error
	switch global.Config.Mysql.LogLevel {
	case "info":
		logLevel = logger.Info
	case "warn":
		logLevel = logger.Warn
	}
	mysqlLogger := logger.Default.LogMode(logLevel)

	// 连接 Mysql 数据库
	dsn := global.Config.Mysql.Dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		logrus.Fatalln(fmt.Sprintf("[%s] Mysql 连接失败, %s", dsn, err.Error()))
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)                   // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)                  // 最大连接数
	sqlDB.SetConnMaxLifetime(time.Hour * 4)     // 单个连接的最长复用时间
	
	return db 
}