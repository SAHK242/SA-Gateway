package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type CorsMiddleware struct {
}

func (c *CorsMiddleware) ShouldSkip(ctx *fiber.Ctx) bool {
	// TODO implement me
	panic("implement me")
}

func (c *CorsMiddleware) OnIntercept(ctx *fiber.Ctx) error {
	// TODO implement me
	panic("implement me")
}

func (c *CorsMiddleware) AsFiberMiddleware() fiber.Handler {
	// TODO: Config allow origins
	return cors.New(cors.Config{
		Next: func(c *fiber.Ctx) bool {
			return false
		},
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
			fiber.MethodOptions,
		}, ","),
		AllowCredentials: false,
		MaxAge:           60,
	})
}

func NewCorsMiddleware() *CorsMiddleware {
	return &CorsMiddleware{}
}

func (c *CorsMiddleware) Order() int {
	return 1
}
