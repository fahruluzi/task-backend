package service

import (
	"fetch/application/model"
	"fetch/application/repository"
	"strconv"
)

type (
	IFetchService interface {
		FetchDataAndAddUSDCurrency() (response []model.FetchResponse, err error)
	}

	FetchService struct {
		RestEFisheryRepository repository.IRestEFisheryRepository
		RestCurrencyRepository repository.ICacheCurrencyRepository
	}
)

func NewFetchService() IFetchService {
	return &FetchService{
		RestEFisheryRepository: repository.NewRestEFisheryRepository(),
		RestCurrencyRepository: repository.NewCacheCurrencyRepository(),
	}
}

func (service *FetchService) FetchDataAndAddUSDCurrency() (response []model.FetchResponse, err error) {
	idrToUsdCurrency, err := service.RestCurrencyRepository.ConvertIDRToUSD()
	if err != nil {
		return
	}

	eFisheryData, err := service.RestEFisheryRepository.RestEFishery()
	if err != nil {
		return
	}

	for _, value := range eFisheryData {
		priceIDR, err := strconv.Atoi(value.Price)
		if err != nil {
			priceIDR = 0
		}

		response = append(response, model.FetchResponse{
			UUID:        value.UUID,
			Commodity:   value.Commodity,
			Province:    value.Province,
			City:        value.City,
			Size:        value.Size,
			Price:       value.Price,
			PriceUSD:    idrToUsdCurrency.IDRToUSD * float64(priceIDR),
			TimeParsing: value.TimeParsing,
			Timestamp:   value.Timestamp,
		})
	}
	return response, nil
}
