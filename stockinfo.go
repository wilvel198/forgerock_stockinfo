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
