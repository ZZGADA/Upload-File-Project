package main

import (
	"UploadFileProject/src/config"
	"UploadFileProject/src/controller"
	"fmt"
)

func main() {
	config.LoadResource("application.yaml")

	controller.InitController(config.Router)

	var runPort = fmt.Sprintf(":%d", config.ServerAllConfig.Port)

	if err := config.Router.Run(runPort); err != nil {
		config.Log.Panicf("gin start panic,%#v", err)
	}
}
