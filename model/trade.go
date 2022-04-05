package model

// 下单

type TradeOrder struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		ClOrdID string `json:"clOrdId"`
		OrdID   string `json:"ordId"`
		Tag     string `json:"tag"`
		SCode   string `json:"sCode"`
		SMsg    string `json:"sMsg"`
	} `json:"data"`
}

// 取消订单

type TradeCancelOrder struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		ClOrdID string `json:"clOrdId"`
		OrdID   string `json:"ordId"`
		SCode   string `json:"sCode"`
		SMsg    string `json:"sMsg"`
	} `json:"data"`
}

// 订单详情结果

type TradeOrderInfo struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		InstType        string `json:"instType"`
		InstID          string `json:"instId"`
		Ccy             string `json:"ccy"`
		OrdID           string `json:"ordId"`
		ClOrdID         string `json:"clOrdId"`
		Tag             string `json:"tag"`
		Px              string `json:"px"`
		Sz              string `json:"sz"`
		Pnl             string `json:"pnl"`
		OrdType         string `json:"ordType"`
		Side            string `json:"side"`
		PosSide         string `json:"posSide"`
		TdMode          string `json:"tdMode"`
		AccFillSz       string `json:"accFillSz"`
		FillPx          string `json:"fillPx"`
		TradeID         string `json:"tradeId"`
		FillSz          string `json:"fillSz"`
		FillTime        string `json:"fillTime"`
		Source          string `json:"source"`
		State           string `json:"state"`
		AvgPx           string `json:"avgPx"`
		Lever           string `json:"lever"`
		TpTriggerPx     string `json:"tpTriggerPx"`
		TpTriggerPxType string `json:"tpTriggerPxType"`
		TpOrdPx         string `json:"tpOrdPx"`
		SlTriggerPx     string `json:"slTriggerPx"`
		SlTriggerPxType string `json:"slTriggerPxType"`
		SlOrdPx         string `json:"slOrdPx"`
		FeeCcy          string `json:"feeCcy"`
		Fee             string `json:"fee"`
		RebateCcy       string `json:"rebateCcy"`
		Rebate          string `json:"rebate"`
		TgtCcy          string `json:"tgtCcy"`
		Category        string `json:"category"`
		UTime           string `json:"uTime"`
		CTime           string `json:"cTime"`
	} `json:"data"`
}

// 未完成订单请求

type OrdersPendingReq struct {
	InstType string `json:"instType" form:"instType"`
	Uly      string `json:"uly" form:"uly"`
	InstId   string `json:"instId" form:"instId"`
	OrdType  string `json:"ordType" form:"ordType"`
	State    string `json:"state" form:"state"`
	After    string `json:"after" form:"after"`
	Before   string `json:"before" form:"before"`
	Limit    string `json:"limit" form:"limit"`
}

// 未完成订单返回

type OrdersPendingRsp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		AccFillSz       string `json:"accFillSz"`
		AvgPx           string `json:"avgPx"`
		CTime           string `json:"cTime"`
		Category        string `json:"category"`
		Ccy             string `json:"ccy"`
		ClOrdID         string `json:"clOrdId"`
		Fee             string `json:"fee"`
		FeeCcy          string `json:"feeCcy"`
		FillPx          string `json:"fillPx"`
		FillSz          string `json:"fillSz"`
		FillTime        string `json:"fillTime"`
		InstID          string `json:"instId"`
		InstType        string `json:"instType"`
		Lever           string `json:"lever"`
		OrdID           string `json:"ordId"`
		OrdType         string `json:"ordType"`
		Pnl             string `json:"pnl"`
		PosSide         string `json:"posSide"`
		Px              string `json:"px"`
		Rebate          string `json:"rebate"`
		RebateCcy       string `json:"rebateCcy"`
		Side            string `json:"side"`
		SlOrdPx         string `json:"slOrdPx"`
		SlTriggerPx     string `json:"slTriggerPx"`
		SlTriggerPxType string `json:"slTriggerPxType"`
		Source          string `json:"source"`
		State           string `json:"state"`
		Sz              string `json:"sz"`
		Tag             string `json:"tag"`
		TdMode          string `json:"tdMode"`
		TgtCcy          string `json:"tgtCcy"`
		TpOrdPx         string `json:"tpOrdPx"`
		TpTriggerPx     string `json:"tpTriggerPx"`
		TpTriggerPxType string `json:"tpTriggerPxType"`
		TradeID         string `json:"tradeId"`
		UTime           string `json:"uTime"`
	} `json:"data"`
}

// 完成订单请求

type OrderHistoryReq struct {
	InstType string `json:"instType" form:"instType"`
	Uly      string `json:"uly" form:"uly"`
	InstId   string `json:"instId" form:"instId"`
	OrdType  string `json:"ordType" form:"ordType"`
	State    string `json:"state" form:"state"`
	Category string `json:"category" form:"category"`
	After    string `json:"after" form:"after"`
	Before   string `json:"before" form:"before"`
	Limit    string `json:"limit" form:"limit"`
}

// 完成订单返回

type OrderHistoryRsp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		InstType        string `json:"instType"`
		InstID          string `json:"instId"`
		Ccy             string `json:"ccy"`
		OrdID           string `json:"ordId"`
		ClOrdID         string `json:"clOrdId"`
		Tag             string `json:"tag"`
		Px              string `json:"px"`
		Sz              string `json:"sz"`
		OrdType         string `json:"ordType"`
		Side            string `json:"side"`
		PosSide         string `json:"posSide"`
		TdMode          string `json:"tdMode"`
		AccFillSz       string `json:"accFillSz"`
		FillPx          string `json:"fillPx"`
		TradeID         string `json:"tradeId"`
		FillSz          string `json:"fillSz"`
		FillTime        string `json:"fillTime"`
		Source          string `json:"source"`
		State           string `json:"state"`
		AvgPx           string `json:"avgPx"`
		Lever           string `json:"lever"`
		TpTriggerPx     string `json:"tpTriggerPx"`
		TpTriggerPxType string `json:"tpTriggerPxType"`
		TpOrdPx         string `json:"tpOrdPx"`
		SlTriggerPx     string `json:"slTriggerPx"`
		SlTriggerPxType string `json:"slTriggerPxType"`
		SlOrdPx         string `json:"slOrdPx"`
		FeeCcy          string `json:"feeCcy"`
		Fee             string `json:"fee"`
		RebateCcy       string `json:"rebateCcy"`
		Rebate          string `json:"rebate"`
		TgtCcy          string `json:"tgtCcy"`
		Pnl             string `json:"pnl"`
		Category        string `json:"category"`
		UTime           string `json:"uTime"`
		CTime           string `json:"cTime"`
	} `json:"data"`
}
