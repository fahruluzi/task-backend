package controller

import (
	"fetch/application/service"
	"fetch/lib"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type (
	IAggregationController interface {
		Aggregation(ctx *fiber.Ctx) error
	}

	AggregationController struct {
		AggregationService service.IAggregationService
	}
)

func NewAggregationController() IAggregationController {
	return &AggregationController{
		AggregationService: service.NewAggregationService(),
	}
}

func (controller *AggregationController) Aggregation(ctx *fiber.Ctx) error {
	result, err := controller.AggregationService.Aggregation()
	if err != nil {
		return lib.ResponseFormatter(ctx, http.StatusInternalServerError, err.Error(), nil, err)
	}

	return lib.ResponseFormatter(ctx, http.StatusOK, "Success Aggregation Data", result, nil)
}
