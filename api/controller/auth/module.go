package authcontroller

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewAuthController,
	NewDepartmentController,
	NewEmployeeController,
)
