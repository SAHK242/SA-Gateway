package authcontroller

import (
	"encoding/base64"
	authmodel "gateway/api/model/auth"
	basemodel "gateway/api/model/base"
	apiutil "gateway/api/util"
	authgrpcclient "gateway/grpc/client/auth"
	grpcutil "gateway/grpc/util"
	"gateway/proto/auth"
	"gateway/proto/gcommon"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"time"
)

type AuthController struct {
	*authgrpcclient.AuthGrpcClient
	redis *redis.Client
}

func NewAuthController(authgrpcclient *authgrpcclient.AuthGrpcClient, redis *redis.Client) *AuthController {
	return &AuthController{AuthGrpcClient: authgrpcclient, redis: redis}
}

// GenerateLoginSecret
// @Summary     Pre-flight Login
// @Description Pre-flight Login
// @Tags        SA - Auth
// @Accept      json
// @Produce     json
// @Param       _               body     authmodel.GenerateLoginSecretRequest true "Request body"
// @Success     200             {object} authmodel.GenerateLoginSecretResponse
// @Failure     400,401,403,500 {object} basemodel.ApiError
// @Router      /auth/pre-flight [post]
func (c *AuthController) GenerateLoginSecret(ctx *fiber.Ctx) error {
	req := new(authmodel.GenerateLoginSecretRequest)
	if err := ctx.BodyParser(req); err != nil {
		return err
	}

	key := lo.RandomString(32, lo.AlphanumericCharset)

	// Store key in Redis with a TTL of 5 minutes
	err := c.redis.Set(ctx.Context(), req.Data, key, 5*time.Minute).Err()
	if err != nil {
		return apiutil.AsApiError(err)
	}

	encodedKey := base64.StdEncoding.EncodeToString([]byte(key))

	return apiutil.AsServerResponse(ctx, &authmodel.GenerateLoginSecretResponse{
		Code: apiutil.AsSuccessCode(),
		Data: authmodel.LoginSecret{Secret: encodedKey},
	})
}

// Login
// @Summary     Employee Login
// @Description Employee Login
// @Tags        SA - Auth
// @Accept      json
// @Produce     json
// @Param       _ body   auth.LoginRequest true "Login Request"
// @Success     200                          {object} authmodel.LoginResponse
// @Failure     400,401,403,500              {object} basemodel.ApiError
// @Router      /auth/login [post]
// @Security    JWT
func (c *AuthController) Login(ctx *fiber.Ctx) error {
	md := grpcutil.WithAnonymousMetadata()

	req := new(auth.LoginRequest)

	if err := ctx.BodyParser(req); err != nil {
		return apiutil.AsApiError(err)
	}

	res, err := c.AuthGrpcClient.Login(req, md)

	if err != nil {
		return apiutil.AsApiError(err)
	}

	if apiutil.HasGrpcError(res.Error) {
		return apiutil.AsApiError(err)
	}

	return apiutil.AsServerResponse(ctx, &authmodel.LoginResponse{
		Code:     apiutil.AsSuccessCode(),
		Token:    res.Token,
		User:     res.User,
		NextStep: res.NextStep,
	})
}

// ChangePassword
// @Summary     Employee Change Password
// @Description Employee Change Password
// @Tags        SA - Auth
// @Accept      json
// @Produce     json
// @Param       _ body   auth.ChangePasswordRequest true "Change password Request"
// @Success     200                          {object} basemodel.ApiEmptyResponse
// @Failure     400,401,403,500              {object} basemodel.ApiError
// @Router      /auth/change-password [post]
// @Security    JWT
func (c *AuthController) ChangePassword(ctx *fiber.Ctx) error {
	md := grpcutil.WithAnonymousMetadata()

	req := new(auth.ChangePasswordRequest)

	if err := ctx.BodyParser(req); err != nil {
		return apiutil.AsApiError(err)
	}

	res, err := c.AuthGrpcClient.ChangePassword(req, md)

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

// CreateEmployee
// @Summary     Create Employee
// @Description Create Employee
// @Tags        SA - Auth
// @Accept      json
// @Produce     json
// @Param       _ body   auth.CreateEmployeeRequest true "CreateEmployeeRequest Request"
// @Success     200                          {object} basemodel.ApiEmptyResponse
// @Failure     400,401,403,500              {object} basemodel.ApiError
// @Router      /auth/create-employee [post]
// @Security    JWT
func (c *AuthController) CreateEmployee(ctx *fiber.Ctx) error {
	md := grpcutil.WithModuleACLMetadata(ctx, gcommon.Module_MODULE_AUTH)

	req := new(auth.CreateEmployeeRequest)

	if err := ctx.BodyParser(req); err != nil {
		return apiutil.AsApiError(err)
	}

	res, err := c.AuthGrpcClient.CreateEmployee(req, md)

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

func (c *AuthController) CreateDepartment(ctx *fiber.Ctx) error {
	md := grpcutil.WithModuleACLMetadata(ctx, gcommon.Module_MODULE_AUTH)

	req := new(auth.CreateDepartmentRequest)

	if err := ctx.BodyParser(req); err != nil {
		return apiutil.AsApiError(err)
	}

	res, err := c.AuthGrpcClient.CreateDepartment(req, md)

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
// @Tags        SA - Auth
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
func (c *AuthController) ListEmployee(ctx *fiber.Ctx) error {
	md := grpcutil.WithModuleACLMetadata(ctx, gcommon.Module_MODULE_AUTH)

	req := c.buildListRequest(ctx)

	res, err := c.AuthGrpcClient.ListEmployee(req, md)

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

func (c *AuthController) buildListRequest(ctx *fiber.Ctx) *auth.ListEmployeeRequest {
	req := &auth.ListEmployeeRequest{
		Pageable:     apiutil.AsPageable(ctx, "code,asc"),
		DepartmentId: ctx.Query("department_id"),
		Search:       ctx.Query("search"),
		EmployeeId:   ctx.Query("employee_id"),
		EmployeeType: auth.EmployeeType(ctx.QueryInt("employee_type")),
	}
	return req
}

// ListDepartment
// @Summary     List Department
// @Description List Department
// @Tags        SA - Auth
// @Accept      json
// @Produce     json
// @Param       page            query    int    false "Page number"
// @Param       size            query    int    false "Page size"
// @Param       sort            query    string false "Sort field. Default to name,asc" Enums(name, code)
// @Param       paging_ignored  query    bool   false "Ignore pagination and return all data"
// @Param       search          query    string false "Search by name"
// @Success     200             {object} authmodel.ListDepartment
// @Failure     400,401,403,500 {object} basemodel.ApiError
// @Router      /auth/list-department [get]
// @Security    JWT
func (c *AuthController) ListDepartment(ctx *fiber.Ctx) error {
	md := grpcutil.WithModuleACLMetadata(ctx, gcommon.Module_MODULE_AUTH)

	req := c.buildListDepartmentRequest(ctx)

	if res, err := c.AuthGrpcClient.ListDepartment(req, md); err != nil {
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

func (c *AuthController) buildListDepartmentRequest(ctx *fiber.Ctx) *auth.ListDepartmentRequest {
	req := &auth.ListDepartmentRequest{
		Pageable: apiutil.AsPageable(ctx, "name,asc"),
		Search:   ctx.Query("search"),
	}
	return req
}
