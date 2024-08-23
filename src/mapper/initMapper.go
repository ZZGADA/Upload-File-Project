package mapper

import (
	"UploadFileProject/src/global"
	"gorm.io/gorm"
)

var mysqlClient *gorm.DB

// 初始化Mapper层
func InitMapper() {
	mysqlClient = global.MySQLClient
}
