package service

import (
	"fmt"
	"strings"
)

type AccountSvr struct {
	Okex
}

/**
   查看账号余额
   @ccy  币种，如 BTC 支持多币种查询（不超过20个），币种之间半角逗号分隔
 */
func (a *AccountSvr) GetBalance(ccys []string)  {
	params := make(map[string]string)
	if len(ccys) > 0 {
		params["ccy"] = strings.Join(ccys, ",")
	}
	res, err := a.SendGetReq("/api/v5/account/balance", params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(res))
	}
}


