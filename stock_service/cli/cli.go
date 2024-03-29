package cli

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/rysmaadit/finantier_test/stock_service/app"
	"github.com/rysmaadit/finantier_test/stock_service/router"
	"github.com/rysmaadit/finantier_test/stock_service/service"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type Cli struct {
	Args []string
}

func NewCli(args []string) *Cli {
	return &Cli{
		Args: args,
	}
}

func (c *Cli) Run(application *app.Application) {
	log.SetLevel(log.InfoLevel)
	log.StandardLogger()
	log.SetOutput(os.Stdout)

	if strings.ToLower(application.Config.LogLevel) == log.DebugLevel.String() {
		log.SetLevel(log.DebugLevel)
	}

	log.SetReportCaller(true)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", application.Config.AppPort),
		Handler: router.NewRouter(service.InstantiateDependencies(application)),
	}

	log.Println(fmt.Sprintf("starting application { %v } on port :%v", application.Config.AppName, application.Config.AppPort))

	go listenAndServe(srv)
	waitForShutdown(srv)
}

func listenAndServe(apiServer *http.Server) {
	err := apiServer.ListenAndServe()

	if err != nil {
		log.WithField("error", err.Error()).Fatal("unable to serve")
	}
}

func waitForShutdown(apiServer *http.Server) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGINT,
		syscall.SIGTERM)

	_ = <-sig

	log.Warn("shutting down")

	if err := apiServer.Shutdown(context.Background()); err != nil {
		log.Println(err)
	}

	log.Warn("shutdown complete")
}
