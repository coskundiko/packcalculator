package server

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"packcalculator/app/handler"
	"packcalculator/config"
	"packcalculator/config/constants/messages"
	"syscall"
)

type Server struct {
	CFG     *config.Config
	Echo    *echo.Echo
	Handler handler.Handler
}

// New sets up the config, Echo, and handler
func New() *Server {

	return &Server{
		CFG:     config.LoadConfig(),
		Echo:    echo.New(),
		Handler: handler.New(),
	}
}
func (s *Server) StartServer() {
	log.Info().Msg(messages.EchoServiceStartingMsg)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	s.Echo.Debug = true

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// We set up the HTTP error handler; for now, we are keeping the default one.
	s.Echo.HTTPErrorHandler = s.Echo.DefaultHTTPErrorHandler
	s.Echo.Use(
		middleware.Recover(),
		middleware.Gzip(),
	)

	// Loads routes
	s.API()

	// We start the HTTP server in a goroutine,allowing the main program to run and listen for exit signals.
	go func() {
		// We use net/http to listen on the configured port and set Echo as the handler.
		serverCnfg := &http.Server{
			Addr:    ":" + s.CFG.Port,
			Handler: s.Echo,
		}

		// We start the listener
		err := serverCnfg.ListenAndServe()
		if err != nil {
			fmt.Print(err)
			cancel()
		}
		serverCnfg.SetKeepAlivesEnabled(true)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	signal.Notify(quit, os.Interrupt, syscall.SIGHUP)
	signal.Notify(quit, os.Interrupt, syscall.SIGQUIT)
	signal.Notify(quit, os.Interrupt, syscall.SIGKILL)
	select {
	case v := <-quit:
		fmt.Printf("signal.Notify: %v \n", v)
	case done := <-ctx.Done():
		fmt.Printf("cts.Done: %v \n", done)
	}
	if err := s.Echo.Shutdown(ctx); err != nil {
		fmt.Printf("Router.Shutdown: %v \n", err)
	}
	fmt.Println("Server Exited Properly")
}
