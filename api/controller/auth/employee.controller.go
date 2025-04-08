package authcontroller

import (
	authmodel "gateway/api/model/auth"
	basemodel "gateway/api/model/base"
	apiutil "gateway/api/util"
	authgrpcclient "gateway/grpc/client/auth"
	grpcutil "gateway/grpc/util"
	"gateway/proto/auth"
	"gateway/proto/gcommon"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type EmployeeController struct {
	*authgrpcclient.EmployeeGrpcClient
	redis *redis.Client
}

func NewEmployeeController(employeegrpcclient *authgrpcclient.EmployeeGrpcClient, redis *redis.Client) *EmployeeController {
	return &EmployeeController{EmployeeGrpcClient: employeegrpcclient, redis: redis}
}

// CreateEmployee
// @Summary     Create Employee
// @Description Create Employee
// @Tags        SA - Employee
// @Accept      json
// @Produce     json
// @Param       _ body   auth.CreateEmployeeRequest true "CreateEmployeeRequest Request"
// @Success     200                          {object} basemodel.ApiEmptyResponse
// @Failure     400,401,403,500              {object} basemodel.ApiError
// @Router      /auth/create-employee [post]
// @Security    JWT
func (c *EmployeeController) CreateEmployee(ctx *fiber.Ctx) error {
	md := grpcutil.WithModuleACLMetadata(ctx, gcommon.Module_MODULE_AUTH)

	req := new(auth.CreateEmployeeRequest)

	if err := ctx.BodyParser(req); err != nil {
		return apiutil.AsApiError(err)
	}

	res, err := c.EmployeeGrpcClient.CreateEmployee(req, md)

	if err != nil {
		return apiutil.AsApiError(err)
	}

	if apiutil.HasGrpcError(res.Error) {
		return apiutil.AsApiError(err)
	}

	return apiutil.AsServerResponse(ctx, &basemodel.ApiEmptyResponse{
		Code:    apiutil.AsSuccessCode(),
		Message: "OK",
	})
}

// ListEmployee
// @Summary     List Employee
// @Description List Employee
// @Tags        SA - Employee
// @Accept      json
// @Produce     json
// @Param       page            query    int    false "Page number"
// @Param       size            query    int    false "Page size"
// @Param       sort            query    string false "Sort field. Default to code,asc" Enums(name, code)
// @Param       paging_ignored  query    bool   false "Ignore pagination and return all data"
// @Param       department_id   query    string false "Filter by department id"
// @Param       search          query    string false "Search by name or email, code, phone number"
// @Param       employee_id    query    string false "Filter by employee id"
// @Param       employee_type  query    int   false "Filter by employee type 1: Doctor, 2: Nurse"
// @Success     200             {object} authmodel.ListEmployee
// @Failure     400,401,403,500 {object} basemodel.ApiError
// @Router      /auth/list-employee [get]
// @Security    JWT
func (c *EmployeeController) ListEmployee(ctx *fiber.Ctx) error {
	md := grpcutil.WithModuleACLMetadata(ctx, gcommon.Module_MODULE_AUTH)

	req := c.buildListRequest(ctx)

	res, err := c.EmployeeGrpcClient.ListEmployee(req, md)

	if err != nil {
		return apiutil.AsApiError(err)
	}

	if apiutil.HasGrpcError(res.Error) {
		return apiutil.AsApiError(err)
	}

	return apiutil.AsServerResponse(ctx, &authmodel.ListEmployee{
		Code: apiutil.AsSuccessCode(),
		Data: authmodel.Employees{
			PageMetadata: apiutil.AsPageMetadata(res.PageMetadata, req.Pageable, res.Employees),
			Employees:    res.Employees,
		},
	})

}

func (c *EmployeeController) buildListRequest(ctx *fiber.Ctx) *auth.ListEmployeeRequest {
	req := &auth.ListEmployeeRequest{
		Pageable:     apiutil.AsPageable(ctx, "code,asc"),
		DepartmentId: ctx.Query("department_id"),
		Search:       ctx.Query("search"),
		EmployeeId:   ctx.Query("employee_id"),
		EmployeeType: auth.EmployeeType(ctx.QueryInt("employee_type")),
	}
	return req
}
