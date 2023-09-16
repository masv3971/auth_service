package main

import (
	"auth_service/internal/apiv1"
	"auth_service/internal/httpserver"
	"auth_service/pkg/configuration"
	"auth_service/pkg/logger"
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type service interface {
	Close(ctx context.Context) error
}

func main() {
	wg := &sync.WaitGroup{}
	ctx := context.Background()

	services := make(map[string]service)

	cfg, err := configuration.Parse(logger.NewSimple("Configuration"))
	if err != nil {
		panic(err)
	}

	log, err := logger.New("vc_issuer", cfg.Log.FolderPath, cfg.Production)
	if err != nil {
		panic(err)
	}

	apiv1Client, err := apiv1.New(ctx, cfg, log.New("apiv1"))
	if err != nil {
		panic(err)
	}
	httpService, err := httpserver.New(ctx, cfg, apiv1Client, log.New("httpserver"))
	services["httpService"] = httpService
	if err != nil {
		panic(err)
	}

	// Handle sigterm and await termChan signal
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	<-termChan // Blocks here until interrupted

	mainLog := log.New("main")
	mainLog.Info("HALTING SIGNAL!")

	for serviceName, service := range services {
		if err := service.Close(ctx); err != nil {
			mainLog.Trace("serviceName", serviceName, "error", err)
		}
	}

	wg.Wait() // Block here until are workers are done

	mainLog.Info("Stopped")
}
