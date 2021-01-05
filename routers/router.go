package routers

import (
	_ "Gin-Template/docs" //这个很重要,把example替换成你的项目目录
	"Gin-Template/middleware/template"
	"Gin-Template/pkg/setting"
	v1 "Gin-Template/routers/api/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)

	// swagger
	if setting.ServerSetting.RunMode != gin.ReleaseMode {
		url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}

	// apis
	apiv1 := r.Group("/api/v1")
	apiv1.Use(template.MiddlewareTemplate()) // token verify
	{
		apiv1.GET("/template", v1.Template)
	}

	return r
}
