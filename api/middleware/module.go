package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type MyMiddleware interface {
	// ShouldSkip returns true if the middleware should be skipped
	ShouldSkip(ctx *fiber.Ctx) bool
	// OnIntercept is called when the middleware is intercepted
	OnIntercept(ctx *fiber.Ctx) error
	// AsFiberMiddleware returns the middleware as a Fiber middleware
	AsFiberMiddleware() fiber.Handler
	// Order returns the order of the middleware. Default is 0 meaning the middleware is executed in arbitrary order
	Order() int
}

func AsFiberMiddleware(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(MyMiddleware)),
		fx.ResultTags(`group:"middlewares"`),
	)
}

var Module = fx.Provide(
	AsFiberMiddleware(NewAuthMiddleware),
	AsFiberMiddleware(NewCorsMiddleware),
)
