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

type DepartmentController struct {
	*authgrpcclient.DepartmentGrpcClient
	redis *redis.Client
}

func NewDepartmentController(departmentgrpcclient *authgrpcclient.DepartmentGrpcClient, redis *redis.Client) *DepartmentController {
	return &DepartmentController{DepartmentGrpcClient: departmentgrpcclient, redis: redis}
}

// CreateDepartment
// @Summary     Create Department
// @Description Create Department
// @Tags        SA - Department
// @Accept      json
// @Produce     json
// @Param       _ body   auth.CreateDepartmentRequest true "CreateDepartmentRequest Request"
// @Success     200                          {object} basemodel.ApiEmptyResponse
// @Failure     400,401,403,500              {object} basemodel.ApiError
// @Router      /auth/department/create-department [post]
// @Security    JWT
func (c *DepartmentController) CreateDepartment(ctx *fiber.Ctx) error {
	md := grpcutil.WithModuleACLMetadata(ctx, gcommon.Module_MODULE_AUTH)

	req := new(auth.CreateDepartmentRequest)

	if err := ctx.BodyParser(req); err != nil {
		return apiutil.AsApiError(err)
	}

	res, err := c.DepartmentGrpcClient.CreateDepartment(req, md)

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

// ListDepartment
// @Summary     List Department
// @Description List Department
// @Tags        SA - Department
// @Accept      json
// @Produce     json
// @Param       page            query    int    false "Page number"
// @Param       size            query    int    false "Page size"
// @Param       sort            query    string false "Sort field. Default to name,asc" Enums(name, code)
// @Param       paging_ignored  query    bool   false "Ignore pagination and return all data"
// @Param       search          query    string false "Search by name"
// @Success     200             {object} authmodel.ListDepartment
// @Failure     400,401,403,500 {object} basemodel.ApiError
// @Router      /auth/department/list-department [get]
// @Security    JWT
func (c *DepartmentController) ListDepartment(ctx *fiber.Ctx) error {
	md := grpcutil.WithModuleACLMetadata(ctx, gcommon.Module_MODULE_AUTH)

	req := c.buildListDepartmentRequest(ctx)

	if res, err := c.DepartmentGrpcClient.ListDepartment(req, md); err != nil {
		return apiutil.AsApiError(err)
	} else if res != nil && apiutil.HasGrpcError(res.Error) {
		return apiutil.AsGrpcError(res.Error)
	} else {
		return apiutil.AsServerResponse(ctx, &authmodel.ListDepartment{
			Code: apiutil.AsSuccessCode(),
			Data: authmodel.Departments{
				PageMetadata: apiutil.AsPageMetadata(res.PageMetadata, req.Pageable, res.Departments),
				Departments:  res.Departments,
			},
		})
	}
}

func (c *DepartmentController) buildListDepartmentRequest(ctx *fiber.Ctx) *auth.ListDepartmentRequest {
	req := &auth.ListDepartmentRequest{
		Pageable: apiutil.AsPageable(ctx, "name,asc"),
		Search:   ctx.Query("search"),
	}
	return req
}
