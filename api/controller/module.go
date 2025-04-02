package controller

import (
	auth "gateway/api/controller/auth"
	patient "gateway/api/controller/patient"
	"go.uber.org/fx"
)

var Module = fx.Module("controller",
	auth.Module,
	patient.Module,
)
