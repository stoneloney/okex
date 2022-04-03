package model

type AcccountBlance struct {
	Code string `json:"code"`
	Data []struct {
		AdjEq   string `json:"adjEq"`
		Details []struct {
			AvailBal      string `json:"availBal"`
			AvailEq       string `json:"availEq"`
			CashBal       string `json:"cashBal"`
			Ccy           string `json:"ccy"`
			CrossLiab     string `json:"crossLiab"`
			DisEq         string `json:"disEq"`
			Eq            string `json:"eq"`
			EqUsd         string `json:"eqUsd"`
			FrozenBal     string `json:"frozenBal"`
			Interest      string `json:"interest"`
			IsoEq         string `json:"isoEq"`
			IsoLiab       string `json:"isoLiab"`
			IsoUpl        string `json:"isoUpl"`
			Liab          string `json:"liab"`
			MaxLoan       string `json:"maxLoan"`
			MgnRatio      string `json:"mgnRatio"`
			NotionalLever string `json:"notionalLever"`
			OrdFrozen     string `json:"ordFrozen"`
			Twap          string `json:"twap"`
			UTime         string `json:"uTime"`
			Upl           string `json:"upl"`
			UplLiab       string `json:"uplLiab"`
			StgyEq        string `json:"stgyEq"`
		} `json:"details"`
		Imr         string `json:"imr"`
		IsoEq       string `json:"isoEq"`
		MgnRatio    string `json:"mgnRatio"`
		Mmr         string `json:"mmr"`
		NotionalUsd string `json:"notionalUsd"`
		OrdFroz     string `json:"ordFroz"`
		TotalEq     string `json:"totalEq"`
		UTime       string `json:"uTime"`
	} `json:"data"`
	Msg string `json:"msg"`
}


type AcccountPositions struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Adl         string `json:"adl"`
		AvailPos    string `json:"availPos"`
		AvgPx       string `json:"avgPx"`
		CTime       string `json:"cTime"`
		Ccy         string `json:"ccy"`
		DeltaBS     string `json:"deltaBS"`
		DeltaPA     string `json:"deltaPA"`
		GammaBS     string `json:"gammaBS"`
		GammaPA     string `json:"gammaPA"`
		Imr         string `json:"imr"`
		InstID      string `json:"instId"`
		InstType    string `json:"instType"`
		Interest    string `json:"interest"`
		Last        string `json:"last"`
		UsdPx       string `json:"usdPx"`
		Lever       string `json:"lever"`
		Liab        string `json:"liab"`
		LiabCcy     string `json:"liabCcy"`
		LiqPx       string `json:"liqPx"`
		MarkPx      string `json:"markPx"`
		Margin      string `json:"margin"`
		MgnMode     string `json:"mgnMode"`
		MgnRatio    string `json:"mgnRatio"`
		Mmr         string `json:"mmr"`
		NotionalUsd string `json:"notionalUsd"`
		OptVal      string `json:"optVal"`
		PTime       string `json:"pTime"`
		Pos         string `json:"pos"`
		PosCcy      string `json:"posCcy"`
		PosID       string `json:"posId"`
		PosSide     string `json:"posSide"`
		ThetaBS     string `json:"thetaBS"`
		ThetaPA     string `json:"thetaPA"`
		TradeID     string `json:"tradeId"`
		QuoteBal    string `json:"quoteBal"`
		BaseBal     string `json:"baseBal"`
		UTime       string `json:"uTime"`
		Upl         string `json:"upl"`
		UplRatio    string `json:"uplRatio"`
		VegaBS      string `json:"vegaBS"`
		VegaPA      string `json:"vegaPA"`
	} `json:"data"`
}
