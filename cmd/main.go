package main

import (
	"api-layout/internal"
	"api-layout/internal/database"
	"api-layout/internal/module"
	"api-layout/internal/service"
	"api-layout/pkg/logger"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	debugMode = flag.Bool("debugMode", false, "whether to enable debug mode")
)

func init() {
	exePath, err := os.Executable()
	if err != nil {
		return
	}
	out := &lumberjack.Logger{
		Filename:   filepath.Join(filepath.Dir(exePath), "api.log"),
		MaxSize:    50,
		MaxBackups: 20,
		MaxAge:     90,
		Compress:   true,
	}
	logger.SetDefaultLogger(log.LstdFlags|log.Lmicroseconds|log.Llongfile|log.Lmsgprefix, "", out)
}

func main() {
	viper.SetConfigName("api.config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		logger.Error("failed to load config file: %v", err)
		return
	}

	loc, _ := time.LoadLocation(viper.GetString("timezone"))
	time.Local = loc

	localAddr := fmt.Sprintf("%s:%d", viper.GetString("listen.host"), viper.GetInt("listen.port"))

	flag.Parse()
	api := internal.NewApi(localAddr, *debugMode)

	// database
	pg, err := database.NewPostgres(viper.GetString("database.main"))
	if err != nil {
		logger.Error("failed to init postgres: %v", err)
		return
	}
	defer pg.Close()

	rds, err := database.NewRedis(viper.GetString("database.cache"))
	if err != nil {
		logger.Error("failed to init redis: %v", err)
		return
	}
	defer rds.Close()

	// service
	svcExample := service.NewExample(pg, rds)

	// module
	modExample := module.NewExample(svcExample)
	modSwagger := module.NewSwagger()

	api.Install(modExample)
	api.Install(modSwagger)

	if err := api.Run(); err != nil {
		logger.Error("api error: %v", err)
	}
}
