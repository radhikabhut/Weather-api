package router

import (
	"wether-api/pkg/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {

	engine.GET("/getData", handler.GetData)

}
