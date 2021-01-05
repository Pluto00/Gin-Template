package app

import (
	"Gin-Template/models"
	"Gin-Template/pkg/e"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, models.Message{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}
