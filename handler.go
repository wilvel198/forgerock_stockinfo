package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	processJsonString(serviceData)

	//fmt.Println(serviceData)
	return stockDataResults
}

// Service to retrieve stock information from data provider
func retrieveStockInfoSvc() string {
	var stockInformation string
	response, err := http.Get("https://www.alphavantage.co/query?apikey=C227WD9W3LUVKVV9&function=TIME_SERIES_DAILY_ADJUSTED&symbol=IBM")

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

func processJsonString(stockJsonValue string) {
	fmt.Println("------> Processing the data from service <--------")

	byteValue := []byte(stockJsonValue)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println("-------> SUCCESS !!!!! <------")
}
