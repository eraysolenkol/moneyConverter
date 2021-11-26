package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ApiResponse struct {
	Eur EUR
	Gbp GBP
	Jpy JPY
	Aud AUD
	Chf CHF
	Cad CAD
	Pln PLN
	Try TRY
	Cny CNY
}

type EUR struct {
	Name        string
	Rate        float64
	Date        string
	InverseRate float64
}

type GBP struct {
	Name        string
	Rate        float64
	Date        string
	InverseRate float64
}

type JPY struct {
	Name        string
	Rate        float64
	Date        string
	InverseRate float64
}

type AUD struct {
	Name        string
	Rate        float64
	Date        string
	InverseRate float64
}

type CHF struct {
	Name        string
	Rate        float64
	Date        string
	InverseRate float64
}

type CAD struct {
	Name        string
	Rate        float64
	Date        string
	InverseRate float64
}

type PLN struct {
	Name        string
	Rate        float64
	Date        string
	InverseRate float64
}

type TRY struct {
	Name        string
	Rate        float64
	Date        string
	InverseRate float64
}

type CNY struct {
	Name        string
	Rate        float64
	Date        string
	InverseRate float64
}

// This function converts the money and prints it.
func convertMoney(nameOfMoney string, apiResponse ApiResponse) {

	var money float64
	var convertedMoney float64
	var updateDate string

	fmt.Printf("Please enter the USD amount")
	fmt.Scanln(&money)

	if nameOfMoney == apiResponse.Eur.Name {

		convertedMoney = apiResponse.Eur.Rate
		updateDate = apiResponse.Eur.Date

	} else if nameOfMoney == apiResponse.Gbp.Name {

		convertedMoney = apiResponse.Gbp.Rate
		updateDate = apiResponse.Gbp.Date

	} else if nameOfMoney == apiResponse.Jpy.Name {

		convertedMoney = apiResponse.Jpy.Rate
		updateDate = apiResponse.Jpy.Date

	} else if nameOfMoney == apiResponse.Aud.Name {

		convertedMoney = apiResponse.Aud.Rate
		updateDate = apiResponse.Aud.Date

	} else if nameOfMoney == apiResponse.Chf.Name {

		convertedMoney = apiResponse.Chf.Rate
		updateDate = apiResponse.Chf.Date

	} else if nameOfMoney == apiResponse.Cad.Name {

		convertedMoney = apiResponse.Cad.Rate
		updateDate = apiResponse.Cad.Date

	} else if nameOfMoney == apiResponse.Pln.Name {

		convertedMoney = apiResponse.Pln.Rate
		updateDate = apiResponse.Pln.Date

	} else if nameOfMoney == apiResponse.Try.Name {

		convertedMoney = apiResponse.Try.Rate
		updateDate = apiResponse.Try.Date

	} else if nameOfMoney == apiResponse.Cny.Name {

		convertedMoney = apiResponse.Cny.Rate
		updateDate = apiResponse.Cny.Date

	}

	fmt.Printf("%.2f USD = %.2f %s \n", money, convertedMoney*money, nameOfMoney)
	fmt.Println("Update date:", updateDate)

}

func main() {

	var apiResponse ApiResponse

	resp, err := http.Get("http://www.floatrates.com/daily/usd.json")

	if err != nil {
		log.Fatal(err)
	}

	// Reading the json file and formatting it
	data, _ := io.ReadAll(resp.Body)
	json.Unmarshal(data, &apiResponse)

	fmt.Println("1.USD-EURO Exchange \n2.USD-Sterling Exchange \n3.USD-Japanese Yen Exchange")
	fmt.Println("4.USD-Australian Dollar Exchange \n5.USD-Swiss Franc Exchange \n6.USD-Canadian Dollar Exchange")
	fmt.Println("7.USD-Polish Zloty Exchange \n8.USD-Turkish Lira Exchange \n9.USD-Chinese Yuan Exchange")
	fmt.Println("Enter the number of operation you want to choose.")

	var pick int
	fmt.Scanln(&pick)

	switch pick {
	case 1:
		convertMoney("Euro", apiResponse)
	case 2:
		convertMoney("U.K. Pound Sterling", apiResponse)
	case 3:
		convertMoney("Japanese Yen", apiResponse)
	case 4:
		convertMoney("Australian Dollar", apiResponse)
	case 5:
		convertMoney("Swiss Franc", apiResponse)
	case 6:
		convertMoney("Canadian Dollar", apiResponse)
	case 7:
		convertMoney("Polish Zloty", apiResponse)
	case 8:
		convertMoney("Turkish Lira", apiResponse)
	case 9:
		convertMoney("Chinese Yuan", apiResponse)
	}

}
