package server

import (
	repo "github.com/kapralovs/hackers-list/internal/hackers/repository/localstorage"
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
	repo := repo.NewLocalStorage()
	uc := usecase.New(repo)
	handler := delivery.NewHandler(uc)
	handler.RegisterHTTPEndpoints(s.router, uc)
	return s.router.Listen(s.port)
}
