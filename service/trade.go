package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"okex/helper"
	"okex/model"
)

type TradeSvr struct {
	Okex
}

/**
  订单详情
*/

func (a *TradeSvr) OrderInfo(apiParams model.TradeOrderInfoReq) (model.TradeOrderInfo, error) {
	params := make(map[string]string)
	if len(apiParams.InstId) > 0 {
		params["instId"] = apiParams.InstId
	}
	if len(apiParams.OrdId) > 0 {
		params["ordId"] = apiParams.OrdId
	}
	if len(apiParams.ClOrdId) > 0 {
		params["clOrdId"] = apiParams.ClOrdId
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

func (a *TradeSvr) Order(params interface{}) (model.TradeOrder, error) {
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

func (a *TradeSvr) OrdersPending(apiParams model.OrdersPendingReq) (model.OrdersPendingRsp, error) {
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
	if len(apiParams.State) > 0 {
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

/**
  策略委托下单
*/

func (a *TradeSvr) OrderAlgo(apiParams model.TradeOrderAlgoReq) (model.TradeOrderAlgoRsp, error) {
	var res model.TradeOrderAlgoRsp
	if len(apiParams.InstId) == 0 {
		return res, errors.New("instId empty")
	}
	if len(apiParams.TdMode) == 0 {
		return res, errors.New("tdMode empty")
	}

	// 保证金模式：isolated：逐仓 ；cross：全仓     非保证金模式：cash：非保证金
	if !helper.InArray(apiParams.TdMode, []string{"isolated", "cross", "cash"}) {
		return res, errors.New("tdMode error")
	}

	// 订单方向 buy：买 sell：卖
	if len(apiParams.Side) == 0 {
		return res, errors.New("side empty")
	}
	if !helper.InArray(apiParams.Side, []string{"buy", "sell"}) {
		return res, errors.New("side error")
	}

	// conditional：单向止盈止损  oco：双向止盈止损  trigger：计划委托  move_order_stop：移动止盈止损  iceberg：冰山委托   twap：时间加权委托
	if len(apiParams.OrdType) == 0 {
		return res, errors.New("ordType empty")
	}
	if !helper.InArray(apiParams.OrdType, []string{"conditional", "oco", "trigger", "move_order_stop", "iceberg", "twap"}) {
		return res, errors.New("ordType error")
	}

	// sz:数量
	if len(apiParams.Sz) == 0 {
		return res, errors.New("sz empty")
	}

	switch apiParams.OrdType {
	case "conditional", "oco":
	case "trigger":
	case "move_order_stop":
	case "iceberg":
	case "twap":

	}

	// todo

	return res, nil

}
