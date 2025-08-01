package server

import (
	"context"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"packcalculator/config"
	"packcalculator/config/constants/messages"
	"packcalculator/handler"
	"strconv"
	"sync"
	"syscall"
	"time"
)

type Server struct {
	CFG          *config.Config
	Echo         *echo.Echo
	Handler      handler.Handler
	Uptime       time.Time      `json:"uptime"`
	RequestCount uint64         `json:"request_count"`
	Status       map[string]int `json:"status"`
	Mutex        sync.RWMutex
}

func New() *Server {
	return &Server{
		CFG:     config.LoadConfig(),
		Echo:    echo.New(),
		Handler: handler.New(),
		Uptime:  time.Now(),
		Status:  map[string]int{},
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
		s.SetStats,
		s.Header,
		s.CORSWithConfig(),
		middleware.Recover(),
		middleware.Gzip(),
		s.CookieStore(),
		s.RequestIDWithConfig(),
	)
	s.API()
	go func() {
		serverCnfg := &http.Server{
			Addr: ":1234",
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

func (s *Server) Stop() {
	//s.GET("/", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, World!")
	//})
	//s.Logger.Fatal(s.Start(":1323"))
}

// CORSWithConfig middleware adds a `Server` header to the response.
func (x *Server) CORSWithConfig() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
	})
}

// RequestIDWithConfig middleware adds a `Server` header to the response.
func (x *Server) RequestIDWithConfig() echo.MiddlewareFunc {
	return middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return fmt.Sprintf("%d%d", time.Now().UnixNano(), rand.Intn(100000))
		},
	})
}

func (x *Server) CookieStore() echo.MiddlewareFunc {
	return session.Middleware(sessions.NewCookieStore([]byte("1234")))
}

func (x *Server) Header(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}

func (s *Server) SetStats(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		s.Mutex.Lock()
		defer s.Mutex.Unlock()
		s.RequestCount++
		s.Status[strconv.Itoa(c.Response().Status)]++
		return nil
	}
}
