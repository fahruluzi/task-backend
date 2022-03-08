package model

type (
	EFisheryDataResponse struct {
		UUID        string `json:"uuid"`
		Commodity   string `json:"komoditas"`
		Province    string `json:"area_provinsi"`
		City        string `json:"area_kota"`
		Size        string `json:"size"`
		Price       string `json:"price"`
		TimeParsing string `json:"tgl_parsed"`
		Timestamp   string `json:"timestamp"`
	}

	FetchResponse struct {
		UUID        string  `json:"uuid"`
		Commodity   string  `json:"commodity"`
		Province    string  `json:"province"`
		City        string  `json:"city"`
		Size        string  `json:"size"`
		Price       string  `json:"price"`
		PriceUSD    float64 `json:"price_usd"`
		TimeParsing string  `json:"time_parsing"`
		Timestamp   string  `json:"timestamp"`
	}
)
