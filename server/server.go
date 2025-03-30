package server

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"strings"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"go.uber.org/fx"

	"gateway/api/model/base"
	"gateway/api/route/common"
	_ "gateway/docs"
)

func NewHttpServer(props HttpServerProps) *fiber.App {
	contextPath := props.Config.MustGetString("CONTEXT_PATH")
	profile := props.Config.MustGetString("PROFILE")

	app := fiber.New(fiber.Config{
		Prefork:                 false,
		BodyLimit:               10 * 1024 * 1024,
		Concurrency:             256 * 1024,
		AppName:                 "SA - Gateway",
		ReduceMemoryUsage:       false,
		ErrorHandler:            errorHandler,
		JSONEncoder:             json.Marshal,
		JSONDecoder:             json.Unmarshal,
		EnableTrustedProxyCheck: profile != "local",
	})

	// Only enable swagger in dev mode
	if profile != "prod" {
		app.Get(getRoutePrefix(contextPath, "/swagger/*"), swagger.HandlerDefault)
	}

	swagger.New(swagger.Config{})

	for _, router := range props.Routers {
		prefix := getRoutePrefix(contextPath, router.SubPath())
		router.Register(&common.RouterProps{
			App:    app,
			Prefix: prefix,
		})
	}

	props.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			port := props.Config.MustGetInt("HTTP_PORT")
			addr := fmt.Sprintf(":%d", port)

			go func() {
				if err := app.Listen(addr); err != nil {
					props.Logger.Panicf("failed to start HTTP server: %v", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})

	return app
}

func getRoutePrefix(basePath string, subPath string) string {
	return fmt.Sprintf("%s%s", basePath, subPath)
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	var code int
	var fiberErr *fiber.Error
	var serverErr *basemodel.ApiError

	var resp interface{}

	if errors.As(err, &serverErr) {
		code = grpcToHttpStatusCode(serverErr.Code)
		resp = serverErr

		return ctx.Status(code).JSON(resp)
	}

	if errors.As(err, &fiberErr) {
		code = fiberErr.Code
	} else {
		code = fiber.StatusInternalServerError
	}

	resp = &basemodel.ApiError{
		Code:    codes.Unknown.String(),
		Message: err.Error(),
	}

	return ctx.Status(code).JSON(resp)
}

func grpcToHttpStatusCode(code string) int {
	if strings.HasSuffix(code, "BAD_REQUEST") {
		return fiber.StatusBadRequest
	}

	if strings.HasSuffix(code, "NOT_FOUND") {
		return fiber.StatusNotFound
	}

	if strings.HasSuffix(code, "UNAUTHORIZED") {
		return fiber.StatusUnauthorized
	}

	if strings.HasSuffix(code, "FORBIDDEN") {
		return fiber.StatusForbidden
	}

	if strings.Contains(code, "VALIDATION") {
		return fiber.StatusBadRequest
	}

	return fiber.StatusInternalServerError
}
