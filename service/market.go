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
  获取单个产品行情信息
 */

func (a *MarketSvr) Ticker(apiParams model.TickerReq) (model.TickerRsp, error) {
	var res model.TickerRsp
	params := make(map[string]string)
	if len(apiParams.InstId) == 0 {
		return res, errors.New("instId empty")
	}
	params["instId"] = apiParams.InstId

	data, err := a.SendGetReq("/api/v5/market/ticker", params)
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


/**
  获取所有产品行情信息
*/

func (a *MarketSvr) Tickers(apiParams model.TickersReq) (model.TickersRsp, error) {
	var res model.TickersRsp
	if len(apiParams.InstType) == 0 {
		return res, errors.New("instType empty")
	}

	params := make(map[string]string)
	params["instType"] = apiParams.InstType
	if len(apiParams.Uly) > 0 {
		params["uly"] = apiParams.Uly
	}

	data, err := a.SendGetReq("/api/v5/market/ticker", params)
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

/**
  获取产品K线
 */
func (a *MarketSvr) Candles(apiParams model.CandlesReq) (model.CandlesRsp, error) {
	var res model.CandlesRsp
	if len(apiParams.InstId) == 0 {
		return res, errors.New("instId empty")
	}

	params := make(map[string]string)
	params["instId"] = apiParams.InstId
	if len(apiParams.Bar) > 0 {
		params["bar"] = apiParams.Bar
	}
	if len(apiParams.After) > 0 {
		params["after"] = apiParams.After
	}
	if len(apiParams.Before) > 0 {
		params["before"] = apiParams.Before
	}
	if len(apiParams.Limit) > 0 {
		params["limit"] = apiParams.Limit
	}

	data, err := a.SendGetReq("/api/v5/market/candles", params)
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
