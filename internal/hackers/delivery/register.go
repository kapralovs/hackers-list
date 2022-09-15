package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kapralovs/hackers-list/internal/hackers"
)

func (h *Handler) RegisterHTTPEndpoints(router *fiber.App, uc hackers.Usecase) {
	handler := NewHandler(uc)
	router.Get("/json/hackers", handler.GetHackersList)
}
