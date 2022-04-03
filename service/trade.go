package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"okex/model"
)

type TradeSvr struct {
	Okex
}

/**
  查看账号余额
  @ccy  币种，如 BTC 支持多币种查询（不超过20个），币种之间半角逗号分隔
*/

func (a *TradeSvr) Order(params map[string]interface{}) (model.AcccountBlance, error)  {
	var res model.AcccountBlance
	data, err := a.SendPostReq("/api/v5/trade/order", params)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}
	if res.Code != "0" {
		return res, errors.New(fmt.Sprintf("code:%s, msg:%s", res.Code, res.Msg))
	}

	return res, nil
}

