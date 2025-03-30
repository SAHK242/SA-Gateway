package common

import (
	"github.com/gofiber/fiber/v2"
)

type RouterProps struct {
	App    *fiber.App
	Prefix string
}

type Router interface {
	Register(props *RouterProps)
	SubPath() string
}
