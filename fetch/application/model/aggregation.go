package model

type (
	AggregationResponse struct {
		Year     string          `json:"year"`
		Month    string          `json:"month"`
		Week     string          `json:"week"`
		Province string          `json:"province"`
		Data     int             `json:"total_data"`
		Size     AggregationSize `json:"size"`
		Price    AggregationSize `json:"price"`
	}

	AggregationSize struct {
		Maximal float64 `json:"maximal"`
		Minimal float64 `json:"minimal"`
		Median  float64 `json:"median"`
		Average float64 `json:"average"`
	}
)
