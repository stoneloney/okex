package model

type OrsUsdt struct {
	Code string `json:"code"`
	Data []struct {
		AskPx     string `json:"askPx"`
		AskSz     string `json:"askSz"`
		BidPx     string `json:"bidPx"`
		BidSz     string `json:"bidSz"`
		High24h   string `json:"high24h"`
		InstID    string `json:"instId"`
		InstType  string `json:"instType"`
		Last      string `json:"last"`
		LastSz    string `json:"lastSz"`
		Low24h    string `json:"low24h"`
		Open24h   string `json:"open24h"`
		SodUtc0   string `json:"sodUtc0"`
		SodUtc8   string `json:"sodUtc8"`
		Ts        string `json:"ts"`
		Vol24h    string `json:"vol24h"`
		VolCcy24h string `json:"volCcy24h"`
	} `json:"data"`
	Msg string `json:"msg"`
}
