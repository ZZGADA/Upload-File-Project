package main

import (
	"UploadFileProject/src/config"
	"UploadFileProject/src/global"
	"fmt"
)

func main() {
	config.LoadResource("application.yaml")

	var runPort = fmt.Sprintf(":%d", config.ServerAllConfig.Port)

	if err := config.Router.Run(runPort); err != nil {
		global.Log.Panicf("gin start panic,%#v", err)
	}
}
