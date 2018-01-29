package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/PRAgarawal/eet/eet"
	esql "github.com/PRAgarawal/eet/sql"
	"github.com/go-kit/kit/log"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", envString("HTTP_ADDR", ":8080"), "HTTP listen address")
		dsn      = flag.String("dsn", envString("EET_SERVICE_DSN", ""), "DSN for database connection")
	)
	flag.Parse()

	// Logging domain.
	var logger log.Logger
	{
		logger = log.NewJSONLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	}
	logger.Log("message", "starting...")

	// Business domain.
	var mysqlRepo *esql.Repository
	if *dsn != "" {
		var err error
		mysqlRepo, err = esql.NewRepository("mysql", *dsn)
		if nil != err {
			logger.Log("message", "failed creating MySQL repository", "err", err, "severity", "CRITICAL")
			os.Exit(1)
		}
	}
	var eeService eet.Service
	{
		var eeRepo eet.Repository
		eeRepo = mysqlRepo

		eeService = eet.NewService(eeRepo)
		eeService = eet.LoggingMiddleware(logger)(eeService)
	}

	// Endpoint domain.
	itemEndpoints := eet.MakeEetServerEndpoints(eeService)

	// Mechanical domain.
	errc := make(chan error)
	ctx := context.Background()

	// Interrupt handler.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	// HTTP transport.
	go func() {
		httpLogger := log.With(logger, "transport", "HTTP")

		mux := http.NewServeMux()
		mux.Handle("/items/", eet.MakeEetHTTPHandler(ctx, itemEndpoints, httpLogger))

		httpLogger.Log("addr", *httpAddr)
		errc <- http.ListenAndServe(*httpAddr, mux)
	}()

	//TODO Kick off routine to periodically send lunch notifications

	logger.Log("terminated", <-errc)
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
