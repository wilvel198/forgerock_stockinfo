package main

import (
	"fmt"
	"os"
	"testing"
)

func TestReturnGeeks(t *testing.T) {

	setEnvDev()

	var stockDataResults StockInfo
	var stockTempResult StockInfo

	Jsoninfo := ""

	contents, err := os.ReadFile("testjson.json")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	Jsoninfo = string(contents)

	stockDataResults.AvgClosing = 128.76399879455568
	stockDataResults.Days = 20
	stockDataResults.Name = "IBM"

	want := stockDataResults
	got := processJsonString(Jsoninfo)

	stockTempResult = processJsonString(Jsoninfo)
	fmt.Println(stockTempResult.AvgClosing)

	if want != got {
		t.Errorf("Test failure")
	}

}
