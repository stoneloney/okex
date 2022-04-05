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
  订单详情
 */

func (a *TradeSvr) OrderInfo(instId, ordId, clOrdId string) (model.TradeOrderInfo, error) {
	params := make(map[string]string)
	if len(instId) > 0 {
		params["instId"] = instId
	}
	if len(ordId) > 0 {
		params["ordId"] = ordId
	}
	if len(clOrdId) > 0 {
		params["clOrdId"] = clOrdId
	}

	var res model.TradeOrderInfo
	data, err := a.SendGetReq("/api/v5/trade/order", params)
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
  下单
  @params 下单请求参数
*/

func (a *TradeSvr) Order(params interface{}) (model.TradeOrder, error)  {
	var res model.TradeOrder
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

/**
  撤销单
 */

func (a *TradeSvr) CancelOrder(params interface{}) (model.TradeCancelOrder, error) {
	var res model.TradeCancelOrder
	data, err := a.SendPostReq("/api/v5/trade/cancel-order", params)
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
   未成交订单列表
 */

func (a *TradeSvr) OrdersPending(apiParams model.OrdersPendingReq) (model.OrdersPendingRsp, error){
	params := make(map[string]string)
	if len(apiParams.InstType) > 0 {
		params["instType"] = apiParams.InstType
	}
	if len(apiParams.Uly) > 0 {
		params["uly"] = apiParams.Uly
	}
	if len(apiParams.InstId) > 0 {
		params["instId"] = apiParams.InstId
	}
	if len(apiParams.OrdType) > 0 {
		params["ordType"] = apiParams.OrdType
	}
	if len(apiParams.State) > 0 {
		params["state"] = apiParams.State
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

	var res model.OrdersPendingRsp
	data, err := a.SendGetReq("/api/v5/trade/orders-pending", params)
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
  历史订单请求 (7天)
 */

func (a *TradeSvr) OrdersHistory(apiParams model.OrderHistoryReq) (model.OrderHistoryRsp, error) {
	var res model.OrderHistoryRsp
	params := make(map[string]string)
	if len(apiParams.InstType) == 0 {
		return res, errors.New("instType empty")
	}
	params["instType"] = apiParams.InstType

	if len(apiParams.Uly) > 0 {
		params["uly"] = apiParams.Uly
	}
	if len(apiParams.InstId) > 0 {
		params["instId"] = apiParams.InstId
	}
	if len(apiParams.OrdType) > 0 {
		params["ordType"] = apiParams.OrdType
	}
	if len(apiParams.State)  > 0 {
		params["state"] = apiParams.State
	}
	if len(apiParams.Category) > 0 {
		params["category"] = apiParams.Category
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

	data, err := a.SendGetReq("/api/v5/trade/orders-history", params)
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