package authroute

import (
	authcontroller "gateway/api/controller/auth"
	"gateway/api/route/common"
)

type AuthRoute struct {
	AuthController *authcontroller.AuthController
}

func NewAuthRoute(authController *authcontroller.AuthController) *AuthRoute {
	return &AuthRoute{AuthController: authController}
}

func (r *AuthRoute) Register(props *common.RouterProps) {
	router := props.App.Group(props.Prefix)

	router.Post("/login", r.AuthController.Login)
	router.Post("/pre-flight", r.AuthController.GenerateLoginSecret)
	router.Post("/change-password", r.AuthController.ChangePassword)
}

func (r *AuthRoute) SubPath() string {
	return "/auth"
}
