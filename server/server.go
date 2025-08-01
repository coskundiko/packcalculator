package server

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
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

func New() *Server {
	return &Server{
		CFG:     config.LoadConfig(),
		Echo:    echo.New(),
		Handler: handler.New(),
	}
}
func (s *Server) StartServer() {
	log.Info().Msg(messages.AfzaServiceStartingMsg)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	s.Echo.Debug = true
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.LevelFieldName = "lvl"
	zerolog.TimestampFieldName = "time"
	zerolog.MessageFieldName = "msg"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s.Echo.HTTPErrorHandler = s.Echo.DefaultHTTPErrorHandler
	s.Echo.Use(
		middleware.Recover(),
		middleware.Gzip(),
	)
	s.API()
	go func() {
		serverCnfg := &http.Server{
			Addr: ":" + s.CFG.Port,
			//Addr:    ":" + s.CFG.APP.Port,
			Handler: s.Echo,
		}
		err := serverCnfg.ListenAndServe()
		if err != nil {
			log.Print(err)
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
