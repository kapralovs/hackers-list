package delivery

import (
	"encoding/json"
	"log"

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
	hackers, err := h.usecase.GetHackersList()
	if err != nil {
		log.Fatal(err)
	}
	json, err := json.Marshal(hackers)
	if err != nil {
		log.Fatal(err)
	}
	return c.SendString(string(json))
}
