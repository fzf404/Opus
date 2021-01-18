package main

import (
	"Opus/config"
	"Opus/database"
	"Opus/route"

	"github.com/gin-gonic/gin"

	"github.com/spf13/viper"
)

func main() {
	// 初始化
	config.InitConfig()
	database.InitSQL()

	r := gin.Default()
	r = route.CollectRoute(r)
	port := viper.GetString("commin.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
	
}
