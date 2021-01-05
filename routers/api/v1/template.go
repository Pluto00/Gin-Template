package v1

import (
	"Gin-Template/pkg/app"
	"Gin-Template/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Auth godoc
// @Summary Router Template
// @Description 路由模板
// @Tags Template
// @Accept  application/json
// @Produce  application/json
// @Success 200 {object} models.Message
// @Router /template [get]
func Template(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
