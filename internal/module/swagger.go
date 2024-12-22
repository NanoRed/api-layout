package module

import (
	_ "api-layout/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Swagger struct {
}

func NewSwagger() *Swagger {
	swagger := &Swagger{}
	return swagger
}

func (s *Swagger) Attach(router *gin.Engine) {
	swagger := router.Group("/api/module/swagger")
	{
		swagger.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
