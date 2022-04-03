package account

import (
	"github.com/gin-gonic/gin"
	"okex/api"
	"okex/service"
)

// ======= 查看持仓 ======

func PositionsHttp(c *gin.Context) {
	api.DoHttpProcess(new(PositionsApi), c)
}

type PositionsApi struct {
	api.Base
}

/**
   instType String	否	产品类型
				MARGIN：币币杠杆
				SWAP：永续合约
				FUTURES：交割合约
                OPTION：期权
			    instType和instId同时传入的时候会校验instId与instType是否一致，结果返回instId的持仓信息
	instId	String	否	交易产品ID，如：BTC-USD-190927-5000-C
			支持多个instId查询（不超过10个），半角逗号分隔
	posId	String	否	持仓ID
			支持多个posId查询（不超过20个），半角逗号分割

    https://www.okx.com/docs-v5/zh/#rest-api-account-get-positions
 */

func (a *PositionsApi) ProcessHttp() {
	instType := a.Ctx.DefaultQuery("instType", "")
	instId := a.Ctx.DefaultQuery("instId", "")
	posId := a.Ctx.DefaultQuery("posId", "")

	res, err := new(service.AccountSvr).GetPositions(instType, instId, posId)
	if err != nil {
		a.Response(3001, "", "查询错误", err.Error())
		return
	}

	a.Response(0, res, "success", "")
}
