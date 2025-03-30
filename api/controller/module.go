package controller

import (
	auth "gateway/api/controller/auth"
	"go.uber.org/fx"
)

var Module = fx.Module("controller",
	auth.Module,
)
