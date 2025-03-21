package main

import (
	"wether-api/pkg/redis"
	"wether-api/pkg/resty"
	"wether-api/pkg/router"

	"github.com/gin-gonic/gin"
)

func main() {
	redis.NewRedisClient("localhost:6379","",0)
	resty.NewRestClient()
	engine := gin.New()

	router.InitRouter(engine)

	engine.Run()
}
