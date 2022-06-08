package model

type StrategyLog struct {
	Name       string  `json:"name"`
	Number     float64 `json:"number"`
	Amount     float64 `json:"amount"`
	SetPrice   float64 `json:"set_price"`
	FinalPrice float64 `json:"final_price"`
	FinalType  int     `json:"final_type"`
	FinalDate  string  `json:"final_date"`
}
