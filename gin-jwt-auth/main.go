package main

import (
	"fmt"
	"gin-jwt-auth/api/router"
	"gin-jwt-auth/config"
	"gin-jwt-auth/db/initializers"
	"gin-jwt-auth/logger"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	logger.InitLogger()
	defer logger.Log.Sync()

	initializers.ConnectDB()
	initializers.ConnectRedis()
}

func main() {
	fmt.Println("Hello auth")
	r := gin.Default()
	router.GetRoute(r)

	r.Run()
}
