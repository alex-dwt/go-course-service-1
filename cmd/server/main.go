package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"alex/test/internal/model"
	"alex/test/internal/repository/postgresql"
	"alex/test/internal/service/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"alex/test/internal/config"
	"alex/test/internal/http/simple_server"
	"go.uber.org/zap"
)

func main() {
	envCfg, err := config.NewEnvConfig()
	if err != nil {
		log.Fatalf("failed to read env-config: %v", err)
	}

	yamlCfg, err := config.NewYamlConfig(envCfg.YamlConfigPath)
	if err != nil {
		log.Fatalf("failed to read yaml-config: %v", err)
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

	//logger.Debug("DEBUG-DEBUG-DEBUG-DEBUG")
	//logger.Info(
	//	"config loaded",
	//	zap.String("env-config", fmt.Sprintf("%+v", envCfg)),
	//)
	//logger.Info(
	//	"yaml loaded",
	//	zap.String("env-config", fmt.Sprintf("%+v", yamlCfg)),
	//)

	// DB
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		yamlCfg.DBCfg.Host,
		yamlCfg.DBCfg.User,
		yamlCfg.DBCfg.Password,
		yamlCfg.DBCfg.DBName,
		yamlCfg.DBCfg.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("open DB", zap.Error(err))
	}
	if err := db.AutoMigrate(&model.User{}); err != nil {
		logger.Fatal("migrate DB", zap.Error(err))
	}

	userRepository := postgresql.NewUserRepository(db)
	userService := user.New(userRepository)

	simpleServer := simple_server.New(envCfg.ServerPort, logger, envCfg.LoggerDebug, userService)

	done := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		if err := simpleServer.Stop(); err != nil {
			logger.Error("stop server", zap.Error(err))
		}

		logger.Info("server stopped successfully")

		close(done)
	}()

	if err = simpleServer.Start(); err != nil {
		logger.Fatal("start server", zap.Error(err))
	}

	<-done

	logger.Info("program exiting")
}
