package grpcutil

import (
	"fmt"
	"gateway/proto/gcommon"
	"github.com/golang-jwt/jwt/v5"
	"strings"

	"google.golang.org/grpc/metadata"

	"github.com/gofiber/fiber/v2"
)

func WithAnonymousMetadata() metadata.MD {
	return metadata.New(make(map[string]string))
}

func WithModuleACLMetadata(ctx *fiber.Ctx, module gcommon.Module) metadata.MD {
	md := metadata.New(make(map[string]string))

	md.Set(fiber.HeaderAuthorization, ctx.Get(fiber.HeaderAuthorization))

	return md
}

func UnsafeSign(claims jwt.Claims, secret string) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	token = addPrefix(token)

	return token, nil
}

func addPrefix(token string) string {
	return fmt.Sprintf("%s%s", PrefixBearer, token)
}

const (
	PrefixBearer = "Bearer "
)

var (
	unsafeParser = jwt.NewParser(jwt.WithoutClaimsValidation())
)

// UnsafeParse is a function to parses the jwt token without validating the signature.
func UnsafeParse(token string) (*jwt.MapClaims, error) {
	if containsPrefix(token) {
		token = removePrefix(token)
	}

	data, _, err := unsafeParser.ParseUnverified(token, new(jwt.MapClaims))

	if err != nil {
		return nil, err
	}

	if claims, ok := data.Claims.(*jwt.MapClaims); ok {
		return claims, nil
	} else {
		return nil, fmt.Errorf("failed to parse claims: %v", err)
	}
}

func containsPrefix(token string) bool {
	return strings.HasPrefix(token, PrefixBearer)
}

func removePrefix(tokenWithPrefix string) string {
	return RemovePrefix(tokenWithPrefix, PrefixBearer)
}

// RemovePrefix removes the prefix from the string if it exists, otherwise it returns the original string
func RemovePrefix(s string, prefix string) string {
	if len(s) >= len(prefix) && s[:len(prefix)] == prefix {
		return s[len(prefix):]
	}
	return s
}
