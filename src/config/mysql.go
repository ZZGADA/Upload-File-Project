package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var MySQLClient *gorm.DB
var MySQLAllConfig = &MySQLConfig{}

// 初始化MySQLClient
func initMySQLClient() {
	const mysqlConnectStr string = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	MySQLAllConfig = &ProjectConfig.DataBases.Mysql

	dsn := fmt.Sprintf(mysqlConnectStr,
		MySQLAllConfig.Username,
		MySQLAllConfig.Password,
		MySQLAllConfig.Ip,
		MySQLAllConfig.Port,
		MySQLAllConfig.Database)

	MySQLClientDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   // 表前缀
			SingularTable: true, // 禁用表名复数
		}})
	if err != nil {
		panic(err)
	}
	//
	sqlDB, _ := MySQLClientDB.DB()
	// SetMaxIdleConnections 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConnections 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	MySQLClient = MySQLClientDB
}
