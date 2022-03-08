package middleware

import (
	"fetch/lib"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
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

func (middleware *JWTMiddleware) IsAdmin(ctx *fiber.Ctx) (err error) {
	claims := ctx.Locals("claims").(jwt.MapClaims)
	if claims["role"] != "admin" {
		return lib.ResponseFormatter(ctx, http.StatusForbidden, "Role not valid to access!", nil, nil)
	}

	ctx.Next()
	return nil
}
