package middleware

import (
	"errors"
	"gateway/config"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strings"

	model "gateway/api/model/base"
	"gateway/proto/gcommon"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	logger    *zap.SugaredLogger
	config    config.Config
	SecretKey string
}

const ContextPrincipalKey = "ContextPrincipal"

func NewAuthMiddleware(
	logger *zap.SugaredLogger,
	config config.Config,
) *AuthMiddleware {
	secret := config.GetString("JWT_SECRET", "")
	if secret == "" {
		logger.Fatal("JWT_SECRET is required")
	}
	return &AuthMiddleware{
		logger:    logger,
		config:    config,
		SecretKey: secret,
	}
}

func (m *AuthMiddleware) OnIntercept(ctx *fiber.Ctx) error {
	if m.hasAuthorizationHeader(ctx) {
		return m.onUserIntercept(ctx)
	} else {
		return &model.ApiError{
			Code:    gcommon.Code_CODE_UNAUTHORIZED.String(),
			Message: fiber.ErrUnauthorized.Error(),
			Details: []*gcommon.ErrorDetail{
				{
					Key:     "principal_type",
					Message: "principal type not found within current context",
				},
			},
		}
	}
}

func (m *AuthMiddleware) AsFiberMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return m.OnIntercept(ctx)
	}
}

func (m *AuthMiddleware) Order() int {
	return 2
}

// onUserIntercept checks the JWT token, extracts claims, and forwards the request
func (m *AuthMiddleware) onUserIntercept(ctx *fiber.Ctx) error {
	// Extract token from Authorization header
	authHeader := ctx.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return fiber.NewError(fiber.StatusUnauthorized, "Missing or invalid Authorization header")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Validate and parse the token
	claims, err := m.validateToken(tokenString)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired token")
	}

	// Set the extracted user claims in the context
	ctx.Locals(ContextPrincipalKey, claims)

	// Forward the request with the original Authorization header
	ctx.Set("Authorization", "Bearer "+tokenString)

	// Continue to the next middleware/handler
	return ctx.Next()
}

func (m *AuthMiddleware) validateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token uses HMAC (HS256)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(m.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// Extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token claims")
}

func (m *AuthMiddleware) hasAuthorizationHeader(ctx *fiber.Ctx) bool {
	return ctx.Get(fiber.HeaderAuthorization) != ""
}

func (m *AuthMiddleware) hasApiKeyHeader(ctx *fiber.Ctx) bool {
	return ctx.Get("X-Api-Key") != ""
}
