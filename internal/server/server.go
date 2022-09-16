package server

import (
	"github.com/go-redis/redis/v9"
	repo "github.com/kapralovs/hackers-list/internal/hackers/repository/redis"
	"github.com/kapralovs/hackers-list/internal/hackers/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/kapralovs/hackers-list/internal/hackers/delivery"
)

type server struct {
	port   string
	router *fiber.App
}

func New(port string, router *fiber.App) *server {
	return &server{
		port:   port,
		router: router,
	}
}

func (s *server) Run() error {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	redisRepo := repo.NewRedisStorage(redisClient)
	uc := usecase.New(redisRepo)
	handler := delivery.NewHandler(uc)
	handler.RegisterHTTPEndpoints(s.router, uc)
	return s.router.Listen(s.port)
}
