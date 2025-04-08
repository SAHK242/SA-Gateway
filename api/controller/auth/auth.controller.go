package authcontroller

import (
	"encoding/base64"
	authmodel "gateway/api/model/auth"
	basemodel "gateway/api/model/base"
	apiutil "gateway/api/util"
	authgrpcclient "gateway/grpc/client/auth"
	grpcutil "gateway/grpc/util"
	"gateway/proto/auth"
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
		Code:        apiutil.AsSuccessCode(),
		Token:       res.Token,
		User:        res.User,
		NextStep:    res.NextStep,
		AccountType: res.AccountType,
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
