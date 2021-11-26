package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type ApiResponse struct {
	Usd USD
	Eur EUR
	Mxn MXN
	Cad CAD
}

type USD struct {
	Satis   string
	Alis    string
	Degisim string
}

type EUR struct {
	Satis   string
	Alis    string
	Degisim string
}

type MXN struct {
	Satis   string
	Alis    string
	Degisim string
}

type CAD struct {
	Satis   string
	Alis    string
	Degisim string
}

// This function converts the money and prints it.
func convertMoney(nameOfMoney string, apiResponse ApiResponse) {
	var money float64
	var convertedMoney float64

	fmt.Printf("Please enter the %s amount", nameOfMoney)
	fmt.Scanln(&money)

	convertedMoney, err := strconv.ParseFloat(apiResponse.Usd.Satis, 64)

	if nameOfMoney == "EURO" {
		convertedMoney, err = strconv.ParseFloat(apiResponse.Eur.Satis, 64)
	} else if nameOfMoney == "USD" {
		convertedMoney, err = strconv.ParseFloat(apiResponse.Usd.Satis, 64)
	} else if nameOfMoney == "MEXICAN USD" {
		convertedMoney, err = strconv.ParseFloat(apiResponse.Mxn.Satis, 64)
	} else if nameOfMoney == "CAD" {
		convertedMoney, err = strconv.ParseFloat(apiResponse.Cad.Satis, 64)
	}
	// Error Handling
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%.2f %s = %.2f TL", money, nameOfMoney, convertedMoney*money)

}

func main() {

	var apiResponse ApiResponse

	// Getting live exchange rates from an api
	response, err := http.Get("https://api.genelpara.com/embed/doviz.json")

	// Error Handling
	if err != nil {
		log.Fatal(err)
	}

	// Reading the json file and formatting it
	data, _ := io.ReadAll(response.Body)
	json.Unmarshal(data, &apiResponse)

	fmt.Println("1.EURO-TL Exchange \n2.USD-TL Exchange \n3.MEXICAN USD-TL Exchange")
	fmt.Println("4.CAD-TL Exchange")
	fmt.Println("Enter the number of operation you want to choose.")

	var pick int
	fmt.Scanln(&pick)

	// Converting the money on user's pick
	if pick == 1 {
		convertMoney("EURO", apiResponse)
	} else if pick == 2 {
		convertMoney("USD", apiResponse)
	} else if pick == 3 {
		convertMoney("MEXICAN USD", apiResponse)
	} else if pick == 4 {
		convertMoney("CAD", apiResponse)
	}

}
