package route

import (
	authroute "gateway/api/route/auth"
	"gateway/api/route/common"
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
)
