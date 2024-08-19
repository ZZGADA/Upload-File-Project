package main

import (
	"UploadFileProject/src/config"
	"UploadFileProject/src/test"
	"fmt"
)

func tmepMysql() {

}

func main() {
	config.LoadResource("application.yaml")
	test.TestGorm()

	config.Router.Run(fmt.Sprintf(":%d", config.ServerAllConfig.Port))
}
