package middleware

import (
	"fetch/lib"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

type JWTMiddleware struct {
	Secret string
}

func NewJWTMiddleware(secret string) *JWTMiddleware {
	return &JWTMiddleware{
		Secret: secret,
	}
}

func (middleware *JWTMiddleware) Validate(ctx *fiber.Ctx) (err error) {
	// check header
	header := ctx.Get("authorization")
	if !strings.Contains(header, "Bearer") {
		return lib.ResponseFormatter(ctx, http.StatusUnauthorized, "Authentication Token is Required", nil, nil)
	}

	// check token is valid
	token := strings.Replace(header, "Bearer ", "", -1)
	claims, IsValid := lib.NewJWTLibrary(middleware.Secret).ValidateToken(token)
	if !IsValid {
		return lib.ResponseFormatter(ctx, http.StatusUnauthorized, "Authorization token is invalid", nil, nil)
	}

	ctx.Locals("claims", claims)
	ctx.Next()
	return nil
}
