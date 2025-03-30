package api

import (
	"gateway/api/helper"
	"go.uber.org/fx"

	"gateway/api/controller"
	"gateway/api/mapper"
	"gateway/api/route"
)

var Module = fx.Module("api",
	controller.Module,
	mapper.Module,
	route.Module,
	helper.Module,
)
