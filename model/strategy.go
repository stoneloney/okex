package model

type StrategyLog struct {
	Name               string  `json:"name"`
	Number             float64 `json:"number"`
	Amount             float64 `json:"amount"`
	Percentage         float64 `json:"percentage"`
	PercentageIncrease float64 `json:"percentage_increase"`
	PercentageDrop     float64 `json:"percentage_drop"`
	SetPrice           float64 `json:"set_price"`
	FinalPrice         float64 `json:"final_price"`
	FinalType          int     `json:"final_type"`
	FinalDate          string  `json:"final_date"`
}
