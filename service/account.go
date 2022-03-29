package service

import "fmt"

const (
	AccountBaseApi = "/api/account/v3"
)

type AccountSvr struct {
	Okex
}

func (a *AccountSvr) GetWallet() {
	res, err := a.SendGetReq(AccountBaseApi + "/wallet")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}
}


