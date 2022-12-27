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

// function to verify if the service is running
func ServiceStatus(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: Service Status")
	json.NewEncoder(w).Encode(statusInfo)
}

// function used to retrieve the stock information
func GetStockInfo(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: return stock info")
	json.NewEncoder(w).Encode(processGetStockInfo())
}

// function to process the stock information
func processGetStockInfo() StockInfo {
	fmt.Println("<------ Processing stock information ------->")

	var serviceData string
	serviceData = retrieveStockInfoSvc()
	return processJsonString(serviceData)

}

// Service to retrieve stock information from data provider
func retrieveStockInfoSvc() string {
	var stockInformation string
	var apiKey string
	var stockSymbol string
	var serviceURL string

	apiKey = os.Getenv("FR_Stock_API_Key")
	stockSymbol = os.Getenv("FR_Stock_Symbol")

	serviceURL = "https://www.alphavantage.co/query?apikey=" + apiKey + "&function=TIME_SERIES_DAILY_ADJUSTED&symbol=" + stockSymbol
	response, err := http.Get(serviceURL)

	//fmt.Println("URL Called ----> " + serviceURL)

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
	var stockSymbol string
	var dayValue float64

	stockSymbol = os.Getenv("FR_Stock_Symbol")

	dayValue, _ = strconv.ParseFloat(os.Getenv("FR_Stock_Days"), 64)

	fmt.Println("the number of days in the env variable")
	fmt.Println(dayValue)

	byteValue := []byte(stockJsonValue)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	var tempData = result["Time Series (Daily)"]

	jsonStr, err := json.Marshal(tempData)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	plan := []byte(jsonStr)

	var data map[string]*stockData

	json.Unmarshal(plan, &data)

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

		if countItems < dayValue {
			sum = sum + stockCloseVal
		}
		log.Println("Key:", key, "=>", "Element:", element.FourClose)
	}

	stockDataResults.Name = stockSymbol
	stockDataResults.Days = int(dayValue)
	stockDataResults.AvgClosing = (sum / dayValue)

	log.Println("the sum is " + fmt.Sprintf("%f", sum))
	sum = 0
	countItems = 0
	fmt.Println("-------> SUCCESS !!!!! <------")
	return stockDataResults

}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func setEnvDev() {
	os.Setenv("FR_Stock_API_Key", "xxxxxx")
	os.Setenv("FR_Stock_Symbol", "HD")
	os.Setenv("FR_Stock_Days", "20")
}
