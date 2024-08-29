package main

import (
	"fmt"
)

const USD_EUR = 0.89460
const USD_RUB = 91.63180

var sorcesCurr string
var destinationCurr string
var sumConv float64

func main() {
	
	fmt.Println("### Конвертер валюты ###")
	
	avalCurr := availableCurrencies(0)	
	inputData(0, fmt.Sprintf("Введите исходную валюту (%s):", avalCurr))

	inputData(2, "Введите сумму")

	avalCurr = availableCurrencies(1)
	inputData(1, fmt.Sprintf("Введите валюту назначения (%s):", avalCurr))

	result := convert(sorcesCurr, destinationCurr, sumConv)

	fmt.Printf("Результат - %f %s = %f %s", sumConv, sorcesCurr, result, destinationCurr)
			
}

func inputData(mode int, messageText string) {

	fmt.Println(messageText)
	switch mode {
	case 0:
		for{
			fmt.Scanln(&sorcesCurr)
			if checkInputData(mode, sorcesCurr) {
				break
			}
		}		
	case 1:
		for {
			fmt.Scanln(&destinationCurr)
			if checkInputData(mode, destinationCurr) {
				break
			}
		}		
	case 2:
		for {			
			if _, err := fmt.Scanln(&sumConv); err == nil{
				break				
			} else {
				fmt.Println("Веденные данные не являются числом!")
			}
		}							
	default:
		return
	}
	
}

func checkInputData(mode int, currency string) bool{

	if mode == 0 {
		switch currency {
		case "RUB":
			return true
		case "USD":
			return true
		case "EUR":
			return true
		default:
			fmt.Println("Не верно указана валюта")
			avalCurr := availableCurrencies(0)
			fmt.Printf("Доступно: %s\n", avalCurr)
			return false
		}
	} else if mode == 1 && sorcesCurr == "EUR" {
		switch currency {
		case "USD":
			return true
		case "RUB":
			return true
		default:
			fmt.Println("Не верно указана валюта")
			avalCurr := availableCurrencies(1)
			fmt.Printf("Доступно: %s\n", avalCurr)
			return false
		}
	} else if mode == 1 && sorcesCurr == "USD" {
		switch currency {
		case "EUR":
			return true
		case "RUB":
			return true
		default:
			fmt.Println("Не верно указана валюта")
			avalCurr := availableCurrencies(1)
			fmt.Printf("Доступно: %s\n", avalCurr)
			return false
		}
	} else if mode == 1 && sorcesCurr == "RUB"{
		switch currency {
		case "USD":
			return true
		case "EUR":
			return true
		default:
			fmt.Println("Не верно указана валюта")
			avalCurr := availableCurrencies(1)
			fmt.Printf("Доступно: %s\n", avalCurr)
			return false
		}
	}

	return false
	
}

func availableCurrencies(mode int) string{
	
	switch{
	case mode == 0:
		return "RUB/USD/EUR"
	case mode == 1 && sorcesCurr == "USD":
		return "RUB/EUR"
	case mode == 1 && sorcesCurr == "EUR":
		return "RUB/USD"
	case mode == 1 && sorcesCurr == "RUB":
		return "USD/EUR"
	default:
		return ""	
	}

}

func convert(firstVal string, secVal string, sum float64) float64 {

	var result float64

	eurToRub := USD_RUB / USD_EUR

	switch {
	case firstVal == "USD" && secVal == "RUB":
		result = sum * USD_RUB		
	case firstVal == "RUB" && secVal == "USD":
		result = sum / USD_RUB
	case firstVal == "USD" && secVal == "EUR":
		result = sum * USD_EUR
	case firstVal == "EUR" && secVal == "USD":
		result = sum / USD_EUR
	case firstVal == "EUR" && secVal == "RUB":
		result = sum * eurToRub
	case firstVal == "RUB" && secVal == "EUR":
		result = sum / eurToRub
	default:
		result = 0.0
	}
	return result

}
