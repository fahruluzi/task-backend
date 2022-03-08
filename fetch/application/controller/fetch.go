package controller

import (
	"fetch/application/service"
	"fetch/lib"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type (
	IFetchController interface {
		Fetch(ctx *fiber.Ctx) error
	}

	FetchController struct {
		FetchService service.IFetchService
	}
)

func NewFetchController() IFetchController {
	return &FetchController{
		FetchService: service.NewFetchService(),
	}
}

func (controller *FetchController) Fetch(ctx *fiber.Ctx) error {
	response, err := controller.FetchService.FetchDataAndAddUSDCurrency()
	if err != nil {
		return lib.ResponseFormatter(ctx, http.StatusInternalServerError, err.Error(), nil, err)
	}

	return lib.ResponseFormatter(ctx, http.StatusOK, "Success Fetch Data", response, nil)
}
