package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kapralovs/hackers-list/internal/hackers"
)

func RegisterHTTPEndpoints(router *fiber.App, uc hackers.Usecase) {
	h := NewHandler(uc)

	router.Get("/json/hackers", h.GetHackersList)
	// h.GET("", h.Get)
}
