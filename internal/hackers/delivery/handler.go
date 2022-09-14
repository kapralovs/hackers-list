package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kapralovs/hackers-list/internal/hackers"
)

type Hacker struct {
	Name  string
	Score int
}

type Handler struct {
	usecase hackers.Usecase
}

func NewHandler(usecase hackers.Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) GetHackersList(c *fiber.Ctx) error {
	return c.SendString("hello")
}
