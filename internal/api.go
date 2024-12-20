package internal

import (
	"api-layout/internal/module"
	"api-layout/pkg/logger"

	"github.com/gin-gonic/gin"
)

// @title API Layout
// @version 1.0
// @description API Layout

// @contact.name red
// @contact.email radixholms@gmail.com

// @BasePath /api

type Api struct {
	localAddr string
	router    *gin.Engine
}

func NewApi(localAddr string, debugMode bool) *Api {
	api := &Api{
		localAddr: localAddr,
	}
	if debugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	api.router = gin.New()
	api.router.Use(gin.LoggerWithConfig(gin.LoggerConfig{Output: logger.Default().Writer()}))
	api.router.Use(gin.Recovery())
	return api
}

func (a *Api) Install(mod module.Module) {
	mod.Attach(a.router)
}

func (a *Api) Run() (err error) {
	return a.router.Run(a.localAddr)
}
