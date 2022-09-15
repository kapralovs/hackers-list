package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kapralovs/hackers-list/internal/hackers"
)

type Handler struct {
	usecase hackers.Usecase
}

func NewHandler(uc hackers.Usecase) *Handler {
	return &Handler{
		usecase: uc,
	}
}

func (h *Handler) GetHackersList(c *fiber.Ctx) error {
	return c.SendString("Hello")
}
