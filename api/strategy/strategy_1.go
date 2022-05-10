package strategy

import "fmt"

type StrategyOne struct {
	Strategy
}

func (s *StrategyOne) Init() *StrategyOne {
	s.SetCurrency("ORS-USDT")
	s.SetPrice(0.3)

	return s
}

func (s *StrategyOne) Do() {
	data, err := s.GetTickerInfo()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(data)
}

