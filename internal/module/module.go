package module

import "github.com/gin-gonic/gin"

type Module interface {
	Attach(router *gin.Engine)
}
