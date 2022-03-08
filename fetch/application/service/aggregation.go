package service

import (
	"fetch/application/model"
	"fetch/application/repository"
	"fetch/lib"
	"fmt"
	"github.com/montanaflynn/stats"
	"strconv"
	"strings"
)

type (
	IAggregationService interface {
		Aggregation() (result []model.AggregationResponse, err error)
	}

	AggregationService struct {
		RestEFisheryRepository repository.IRestEFisheryRepository
	}
)

func NewAggregationService() IAggregationService {
	return &AggregationService{
		RestEFisheryRepository: repository.NewRestEFisheryRepository(),
	}
}

func (service *AggregationService) Aggregation() (result []model.AggregationResponse, err error) {
	eFisheryData, err := service.RestEFisheryRepository.RestEFishery()
	if err != nil {
		return
	}

	var group = map[string][]model.EFisheryDataResponse{}

	for _, value := range eFisheryData {
		dateConvert := lib.ParseDate(value.TimeParsing)
		weekGroup := fmt.Sprintf("%d-%d-%d", dateConvert.Year(), int(dateConvert.Month()), int(dateConvert.Weekday()))
		keyGroup := fmt.Sprintf("%s#%s", weekGroup, value.Province)
		group[keyGroup] = append(group[keyGroup], value)
	}

	for i := range group {
		key := strings.Split(i, "#")
		date := strings.Split(key[0], "-")

		var listSize []float64
		var listPrice []float64
		for _, data := range group[i] {
			size, _ := strconv.ParseFloat(data.Size, 64)
			listSize = append(listSize, size)
			price, _ := strconv.ParseFloat(data.Price, 64)
			listPrice = append(listPrice, price)
		}

		medianSize, _ := stats.Median(listSize)
		maxSize, _ := stats.Max(listSize)
		minSize, _ := stats.Min(listSize)
		sumSize := 0.0
		for i := range listSize {
			sumSize += listSize[i]
		}
		averageSize := sumSize / float64(len(listSize))

		medianPrice, _ := stats.Median(listPrice)
		maxPrice, _ := stats.Max(listPrice)
		minPrice, _ := stats.Min(listPrice)
		sumPrice := 0.0
		for i := range listPrice {
			sumPrice += listPrice[i]
		}
		averagePrice := sumPrice / float64(len(listPrice))

		if date[0] == "1" && date[1] == "1" && date[2] == "1" {
			result = append(result, model.AggregationResponse{
				Year:     "",
				Month:    "",
				Week:     "",
				Province: key[1],
				Data:     len(group[i]),
				Size: model.AggregationSize{
					Maximal: maxSize,
					Minimal: minSize,
					Median:  medianSize,
					Average: averageSize,
				},
				Price: model.AggregationSize{
					Maximal: maxPrice,
					Minimal: minPrice,
					Median:  medianPrice,
					Average: averagePrice,
				},
			})
		} else {
			result = append(result, model.AggregationResponse{
				Year:     date[0],
				Month:    date[1],
				Week:     date[2],
				Province: key[1],
				Data:     len(group[i]),
				Size: model.AggregationSize{
					Maximal: maxSize,
					Minimal: minSize,
					Median:  medianSize,
					Average: averageSize,
				},
				Price: model.AggregationSize{
					Maximal: maxPrice,
					Minimal: minPrice,
					Median:  medianPrice,
					Average: averagePrice,
				},
			})
		}
	}

	return
}
