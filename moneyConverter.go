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
func convertMoney(nameOfMoney string, apiResponse ApiResponse, convert int) {

	var money float64
	var convertedMoney float64
	var convertedInverseMoney float64
	var updateDate string

	if convert == 1 {
		fmt.Printf("Please enter the USD amount")
	} else if convert == 2 {
		fmt.Printf("Please enter the %s amount", nameOfMoney)
	}

	fmt.Scanln(&money)

	if nameOfMoney == apiResponse.Eur.Name {

		convertedMoney = apiResponse.Eur.Rate
		convertedInverseMoney = apiResponse.Eur.InverseRate
		updateDate = apiResponse.Eur.Date

	} else if nameOfMoney == apiResponse.Gbp.Name {

		convertedMoney = apiResponse.Gbp.Rate
		convertedInverseMoney = apiResponse.Gbp.InverseRate
		updateDate = apiResponse.Gbp.Date

	} else if nameOfMoney == apiResponse.Jpy.Name {

		convertedMoney = apiResponse.Jpy.Rate
		convertedInverseMoney = apiResponse.Jpy.InverseRate
		updateDate = apiResponse.Jpy.Date

	} else if nameOfMoney == apiResponse.Aud.Name {

		convertedMoney = apiResponse.Aud.Rate
		convertedInverseMoney = apiResponse.Aud.InverseRate
		updateDate = apiResponse.Aud.Date

	} else if nameOfMoney == apiResponse.Chf.Name {

		convertedMoney = apiResponse.Chf.Rate
		convertedInverseMoney = apiResponse.Chf.InverseRate
		updateDate = apiResponse.Chf.Date

	} else if nameOfMoney == apiResponse.Cad.Name {

		convertedMoney = apiResponse.Cad.Rate
		convertedInverseMoney = apiResponse.Cad.InverseRate
		updateDate = apiResponse.Cad.Date

	} else if nameOfMoney == apiResponse.Pln.Name {

		convertedMoney = apiResponse.Pln.Rate
		convertedInverseMoney = apiResponse.Pln.InverseRate
		updateDate = apiResponse.Pln.Date

	} else if nameOfMoney == apiResponse.Try.Name {

		convertedMoney = apiResponse.Try.Rate
		convertedInverseMoney = apiResponse.Try.InverseRate
		updateDate = apiResponse.Try.Date

	} else if nameOfMoney == apiResponse.Cny.Name {

		convertedMoney = apiResponse.Cny.Rate
		convertedInverseMoney = apiResponse.Cny.InverseRate
		updateDate = apiResponse.Cny.Date

	}

	if convert == 1 {
		fmt.Printf("%.2f USD = %.2f %s \n", money, convertedMoney*money, nameOfMoney)
	} else if convert == 2 {
		fmt.Printf("%.2f %s = %.2f USD \n", money, nameOfMoney, convertedInverseMoney*money)
	}

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
	var secondPick int
	fmt.Scanln(&pick)

	fmt.Println("Do you want to convert from USD or do inverse convert? \nEnter 1 for convert from USD otherwise enter 2")
	fmt.Scanln(&secondPick)

	switch pick {
	case 1:
		convertMoney("Euro", apiResponse, secondPick)
	case 2:
		convertMoney("U.K. Pound Sterling", apiResponse, secondPick)
	case 3:
		convertMoney("Japanese Yen", apiResponse, secondPick)
	case 4:
		convertMoney("Australian Dollar", apiResponse, secondPick)
	case 5:
		convertMoney("Swiss Franc", apiResponse, secondPick)
	case 6:
		convertMoney("Canadian Dollar", apiResponse, secondPick)
	case 7:
		convertMoney("Polish Zloty", apiResponse, secondPick)
	case 8:
		convertMoney("Turkish Lira", apiResponse, secondPick)
	case 9:
		convertMoney("Chinese Yuan", apiResponse, secondPick)
	}

}
