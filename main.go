package main

import (
	"github.com/gin-gonic/gin"
	"github.com/imzdong/wehcat-mp/wechat"
)

func main() {
	/**
	https://juejin.cn/post/6844904114707496973
	*/
	router := gin.Default()

	router.GET("/wx", wechat.WXCheckSignature)
	router.POST("/wx", wechat.WXMsgReceive)

	router.Run(":8888")
}
