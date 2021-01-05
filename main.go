package main

import (
	"Gin-Template/models"
	"Gin-Template/pkg/gredis"
	"Gin-Template/pkg/setting"
	"Gin-Template/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	//"syscall"
	//"github.com/fvbock/endless"
)

func init() {
	setting.Setup()
	err := gredis.Setup()
	if err != nil {
		log.Panicf("Fail to connect 'redis': %v", err)
	}
	models.Setup()
}

// @title Gin-Template Restful API
// @version 1.0
// @description Gin Restfule API 开发模板

// @contact.name API Support
// @contact.url https://github.com/Pluto00
// @contact.email 1834733586@qq.com
// @BasePath /api/v1
func main() {
	if setting.ServerSetting.RunMode == gin.ReleaseMode {
		//release 热重启
		//endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
		//endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
		//endless.DefaultMaxHeaderBytes = 1 << 20
		//endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
		//
		//server := endless.NewServer(endPoint, routers.InitRouter())
		//server.BeforeBegin = func(add string) {
		//	log.Printf("[info] start http server listening %s", endPoint)
		//	log.Printf("[info] Actual pid is %d", syscall.Getpid())
		//}
		//err := server.ListenAndServe()
		//if err != nil {
		//	log.Printf("Ser")
		//}
	} else {
		//
		router := routers.InitRouter()
		readTimeout := setting.ServerSetting.ReadTimeout
		writeTimeout := setting.ServerSetting.WriteTimeout
		endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
		maxHeaderBytes := 1 << 20

		server := &http.Server{
			Addr:           endPoint,
			Handler:        router,
			ReadTimeout:    readTimeout,
			WriteTimeout:   writeTimeout,
			MaxHeaderBytes: maxHeaderBytes,
		}

		log.Printf("[info] start http server listening %s", endPoint)

		err := server.ListenAndServe()

		if err != nil {
			log.Printf("Fail to run http server: %v", err)
		}
	}
}
