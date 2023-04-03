package main

import (
	"fmt"
	"myblog/models"
	"myblog/pkg/redis"
	"myblog/routers"
	"myblog/settings"
)

func main() {
	settings.SetUp()
	models.SetUp()
	redis.SetUp()

	r := routers.SetUpRouter()
	err := r.Run(fmt.Sprintf(":%d", settings.ServerConf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
