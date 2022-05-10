package strategy

import (
	"fmt"
	"time"
)

type StrategyOne struct {
	Strategy
}

func (s *StrategyOne) Init() *StrategyOne {
	s.SetCurrency("ORS-USDT")
	s.SetPrice(0.3)

	return s
}
func (s *StrategyOne) Run() {
	// 创建定时器
	ticker := time.NewTicker(time.Second * 2)
	go func() {
		for {
			<-ticker.C
			s.Do()
		}
	}()
}

func (s *StrategyOne) Do() {
	data, err := s.GetTickerInfo()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(data)
}

