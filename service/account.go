package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"okex/model"
	"strings"
)

type AccountSvr struct {
	Okex
}

/**
   查看账号余额
   @ccy  币种，如 BTC 支持多币种查询（不超过20个），币种之间半角逗号分隔
 */

func (a *AccountSvr) GetBalance(ccys []string) (model.AcccountBlance, error)  {
	params := make(map[string]string)
	if len(ccys) > 0 {
		params["ccy"] = strings.Join(ccys, ",")
	}
	var res model.AcccountBlance
	data, err := a.SendGetReq("/api/v5/account/balance", params)
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
   查看持仓
 */

func (a *AccountSvr) GetPositions(instType, instId, posId string) (model.AcccountPositions, error) {
	params := make(map[string]string)
	if len(instType) > 0 {
		params["instType"] = instType
	}
	if len(instId) > 0 {
		params["instId"] = instId
	}
	if len(posId) > 0 {
		params["posId"] = posId
	}

	var res model.AcccountPositions
	data, err := a.SendGetReq("/api/v5/account/positions", params)
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
