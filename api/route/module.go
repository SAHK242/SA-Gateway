package route

import (
	authroute "gateway/api/route/auth"
	"gateway/api/route/common"
	patientroute "gateway/api/route/patient"
	"go.uber.org/fx"
)

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(common.Router)),
		fx.ResultTags(`group:"routers"`),
	)
}

var Module = fx.Provide(
	AsRoute(authroute.NewAuthRoute),
	AsRoute(authroute.NewDepartmentRoute),
	AsRoute(authroute.NewEmployeeRoute),

	AsRoute(patientroute.NewPatientRoute),
	AsRoute(patientroute.NewMedicationRoute),
)
