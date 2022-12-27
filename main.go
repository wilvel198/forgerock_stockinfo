package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//setting the dev environment when necessary to run localy without needing to set
	// environmen variables
	//setEnvDev()
	fmt.Println("Stock service started to run")
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/statusinfo", ServiceStatus).Methods("GET")  // ----> Check the status of the service
	r.HandleFunc("/getstockinfo", GetStockInfo).Methods("GET") // ----> To stock information
	log.Fatal(http.ListenAndServe(":8080", r))

}
