package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/alloykh/pack-calc/internal/packs"
	"github.com/alloykh/pack-calc/internal/query"
	"github.com/alloykh/pack-calc/internal/server"
	stdLog "github.com/alloykh/pack-calc/pkg/log"
	stdlibServer "github.com/alloykh/pack-calc/pkg/server"
)

func main() {

	// config setup
	cfg, configSource, err := readConfig()
	if err != nil {
		log.Fatalf("failed to load config: %s", err)
		return
	}

	// logger setup
	logger := stdLog.NewStdioLogger(os.Stdout, os.Stderr, cfg.Log.Verbosity)
	if cfg.Log.Verbosity != stdLog.Silent {
		logger.Trace("using", configSource)
		logger.Trace("Config: ", fmt.Sprintf("%+v", *cfg))
	}

	ctx, cancel := context.WithCancel(context.Background())

	// handle the graceful shutdown
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-interruptChan
		log.Println("context cancellation has started...")
		cancel()
	}()

	packsCalc := packs.NewPacksCalculator(cfg.Packs)

	resolver := query.NewQueryResolver(logger, packsCalc)

	muxHandler := server.NewHttpHandler(logger, resolver)

	srvTr := stdlibServer.RunServer(logger, "main", cfg.ServerAddr(), muxHandler)
	defer srvTr()

	// run the healthcheck server
	healthShutdown := stdlibServer.RunServer(logger, "health-check", cfg.HealthCheckAddr(), stdlibServer.HealthCheckHandler())
	defer healthShutdown()

	<-ctx.Done()

	log.Println("the application has exited successfully")
}
