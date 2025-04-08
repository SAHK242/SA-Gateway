package authroute

import (
	authcontroller "gateway/api/controller/auth"
	"gateway/api/route/common"
)

type DepartmentRoute struct {
	DepartmentController *authcontroller.DepartmentController
}

func NewDepartmentRoute(authController *authcontroller.DepartmentController) *DepartmentRoute {
	return &DepartmentRoute{DepartmentController: authController}
}

func (r *DepartmentRoute) Register(props *common.RouterProps) {
	router := props.App.Group(props.Prefix)

	router.Post("/create-department", r.DepartmentController.CreateDepartment)
	router.Get("/list-department", r.DepartmentController.ListDepartment)
}

func (r *DepartmentRoute) SubPath() string {
	return "/auth/department"
}
