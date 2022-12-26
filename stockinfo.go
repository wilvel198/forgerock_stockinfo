package main

type StockInfo struct {
	Name       string  `json:"name"`
	AvgClosing float64 `json:"avgclosing"`
	Days       int     `json:"days"`
}

type StatusInfo struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}

type stockData struct {
	OneOpen               string `json:"1. open"`
	TwoHigh               string `json:"2. high"`
	ThreeLow              string `json:"3. low"`
	FourClose             string `json:"4. close"`
	FiveAdjustedClose     string `json:"5. adjusted close"`
	SixVolume             string `json:"6. volume"`
	SevenDividendAmount   string `json:"7. dividend amount"`
	EightSplitCoefficient string `json:"8. split coefficient"`
}
