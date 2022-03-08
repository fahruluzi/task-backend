package controller

import (
	"fetch/lib"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ValidateController struct {
}

func NewValidateController() *ValidateController {
	return &ValidateController{}
}

func (controller *ValidateController) Validate(ctx *fiber.Ctx) error {
	lib.ResponseFormatter(ctx, http.StatusOK, "Success Validate Authentication Token", ctx.Locals("claims"), nil)
	return nil
}
