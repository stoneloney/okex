package service

import "fmt"

const (
	AccountBaseApi = "/api/account"
)

type AccountSvr struct {
	Okex
}

func (a *AccountSvr) GetWallet() {
	res, err := a.SendGetReq(AccountBaseApi + "/v3/wallet")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(res))
	}
}

// 查看账号余额
func (a *AccountSvr) GetBalance() {
	res, err := a.SendGetReq("/api/v5/account/balance")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(res))
	}
}


