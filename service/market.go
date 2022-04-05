package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"okex/model"
)

type MarketSvr struct {
	Okex
}

/**
  获取所有产品行情信息
*/

func (a *MarketSvr) Tickers(apiParams model.TickersReq) (model.TickersRsp, error) {
	params := make(map[string]string)
	if len(apiParams.InstType) > 0 {
		params["instType"] = apiParams.InstType
	}
	if len(apiParams.Uly) > 0 {
		params["uly"] = apiParams.Uly
	}

	var res model.TickersRsp
	data, err := a.SendGetReq("/api/v5/market/tickers", params)
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