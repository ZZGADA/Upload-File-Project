package test

import (
	"UploadFileProject/src/config"
	"UploadFileProject/src/entity/bo"
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
)

func TestGorm() {
	// 定义一个切片
	var users []bo.TransactionStudy
	// 查询主键 in （2，4，5）中的
	config.MySQLClient.Find(&users, []int{2, 4, 5})
	fmt.Println(reflect.TypeOf(users).Kind())
	for _, val := range users {
		fmt.Println(val)
	}

	//var user1 connectToMysql.TransactionStudy
	config.MySQLClient.Find(&users, "age > ? and address = ?", 22, "HN")
	for _, val := range users {
		fmt.Printf("%#v", val)
	}
}

func TestGetApi() {
	// /search?query=gin  路由携带参数
	config.Router.GET("/search", func(c *gin.Context) {
		query := c.Query("query")
		more := c.Query("more")
		config.Log.Info("ioioisoaioa")
		config.Log.Warn("pppppppppppp")
		c.String(200, "Search query is "+query+" more is "+more)
	})

}
