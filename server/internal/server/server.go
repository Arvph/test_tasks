package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arvph/test_tasks/internal/handlers"
	"github.com/arvph/test_tasks/internal/services"
	"github.com/sirupsen/logrus"
)

// Server представляет структуру сервера
type Server struct {
	Core     *http.Server
	Handler  *handlers.Handler
	Services *services.Services
	Port     string
	Addr     string
}

// SetServices...
func (s *Server) SetServices(serv *services.Services) {
	s.Services = serv
}

// ServerStart создает сервер
func ServerStart(log *logrus.Logger, conf *Server) error {
	log.Println("Server is starting")

	// подключение хэндлеров
	conf.Handler = handlers.InitHandler(conf.Services)

	r := routers(conf)

	addr := fmt.Sprintf("%s:%s", conf.Addr, conf.Port)
	conf.Core = &http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		if err := conf.Core.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Failed to start server: %s\n", err)
		}
	}()
	log.Printf("Server is running at %s\n", addr)

	chStop := make(chan os.Signal, 1)
	signal.Notify(chStop, syscall.SIGINT, syscall.SIGTERM)

	<-chStop
	log.Println("Stop signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := conf.Core.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %s\n", err)
		return err
	}

	log.Println("Graceful shutdown")
	return nil
}
