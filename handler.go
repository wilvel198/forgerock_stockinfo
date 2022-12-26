package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

var statusInfo = []StatusInfo{
	{Status: "Ok", Info: "Service is running"},
}

func ServiceStatus(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: Service Status")
	json.NewEncoder(w).Encode(statusInfo)
}

func GetStockInfo(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: return stock info")
	json.NewEncoder(w).Encode(processGetStockInfo())
}

func processGetStockInfo() StockInfo {
	fmt.Println("<------ Processing stock information ------->")
	var stockDataResults StockInfo
	var serviceData string

	stockDataResults.AvgClosing = 65.76
	stockDataResults.Name = "IBM"
	stockDataResults.Days = 20

	serviceData = retrieveStockInfoSvc()
	return processJsonString(serviceData)

	//fmt.Println(serviceData)
	//return stockDataResults
}

// Service to retrieve stock information from data provider
func retrieveStockInfoSvc() string {
	var stockInformation string
	var apiKey string
	var stockSymbol string

	os.Setenv("FR_Stock_API_Key", "C227WD9W3LUVKVV9")
	os.Setenv("FR_Stock_Symbol", "IBM")

	apiKey = os.Getenv("FR_Stock_API_Key")
	stockSymbol = os.Getenv("FR_Stock_Symbol")

	response, err := http.Get("https://www.alphavantage.co/query?apikey=" + apiKey + "&function=TIME_SERIES_DAILY_ADJUSTED&symbol=" + stockSymbol)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(responseData))

	stockInformation = string(responseData)

	return stockInformation

}

func processJsonString(stockJsonValue string) StockInfo {
	fmt.Println("------> Processing the data from service <--------")

	var stockDataResults StockInfo
	stockDataResults.Name = "IBM"

	byteValue := []byte(stockJsonValue)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	var tempData = result["Time Series (Daily)"]
	//var symbolInfo = result["Meta Data"]

	fmt.Println("Time Series Daily info ---->")
	fmt.Println(tempData)

	jsonStr, err := json.Marshal(tempData)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(jsonStr))
	}

	plan := []byte(jsonStr)

	var data map[string]*stockData
	json.Unmarshal(plan, &data)
	//fmt.Printf("Data is %+v\n", data)

	var sum float64 = 0
	var countItems float64 = 0

	for key, element := range data {
		countItems = countItems + 1

		var a = element.FourClose
		stockCloseVal, err := strconv.ParseFloat(a, 32)

		if err != nil {
			fmt.Println("Error during conversion")
			return stockDataResults
		}

		sum = sum + stockCloseVal

		fmt.Println("Key:", key, "=>", "Element:", element.FourClose)
	}

	stockDataResults.Days = int(countItems)
	stockDataResults.AvgClosing = roundFloat((sum / countItems), 2)
	fmt.Println("the sum is " + fmt.Sprintf("%f", sum))
	fmt.Println("-------> SUCCESS !!!!! <------")
	return stockDataResults

}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
