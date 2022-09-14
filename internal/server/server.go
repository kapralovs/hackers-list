package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kapralovs/hackers-list/internal/hackers"
	hkhttp "github.com/kapralovs/hackers-list/internal/hackers/delivery"
)

type App struct {
	httpServer     *http.Server
	hackersUsecase hackers.Usecase
}

func NewApp() *App {
	// db := initDB()
	// hackerRepo := redis.NewRedisRepository
	return &App{}
}

func (a *App) Run(port string) error {
	router := fiber.New()
	hkhttp.RegisterHTTPEndpoints(router, a.hackersUsecase)
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}
