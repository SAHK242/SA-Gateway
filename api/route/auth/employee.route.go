package authroute

import (
	authcontroller "gateway/api/controller/auth"
	"gateway/api/route/common"
)

type EmployeeRoute struct {
	EmployeeController *authcontroller.EmployeeController
}

func NewEmployeeRoute(authController *authcontroller.EmployeeController) *EmployeeRoute {
	return &EmployeeRoute{EmployeeController: authController}
}

func (r *EmployeeRoute) Register(props *common.RouterProps) {
	router := props.App.Group(props.Prefix)

	router.Post("/create-employee", r.EmployeeController.CreateEmployee)
	router.Get("/list-employee", r.EmployeeController.ListEmployee)
}

func (r *EmployeeRoute) SubPath() string {
	return "/auth/employee"
}
