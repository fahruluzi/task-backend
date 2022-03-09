package model

type (
	RestCurrencyResponse struct {
		Data CurrencyUSD `json:"data"`
	}

	CurrencyUSD struct {
		USD Currency `json:"USD"`
	}

	Currency struct {
		Code  string  `json:"code"`
		Value float64 `json:"value"`
	}
)
