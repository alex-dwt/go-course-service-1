package main

import (
	"fmt"
	"log"

	"alex/test/internal/config"
	"alex/test/internal/http/simple_server"
	"go.uber.org/zap"
)

func main() {
	envCfg, err := config.NewEnvConfig()
	if err != nil {
		log.Fatalf("failed to read env-config: %v", err)
	}

	var logger *zap.Logger
	if envCfg.LoggerDebug {
		logger, err = zap.NewDevelopment()
		if err != nil {
			log.Fatalf("failed create dev-logger: %v", err)
		}
	} else {
		logger, err = zap.NewProduction()
		if err != nil {
			log.Fatalf("failed create prod-logger: %v", err)
		}
	}
	defer logger.Sync()

	logger.Debug("DEBUG-DEBUG-DEBUG-DEBUG")
	logger.Info(
		"config loaded",
		zap.String("env-config", fmt.Sprintf("%+v", envCfg)),
	)

	simpleServer := simple_server.New(envCfg.ServerPort, logger)
	if err = simpleServer.Start(); err != nil {
		logger.Error("start server", zap.Error(err))
	}
}
