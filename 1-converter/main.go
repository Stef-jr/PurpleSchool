package main

import (
	"fmt"
)

type currencyMap = map [string]float64

var sorcesCurr string
var destinationCurr string
var sumConv float64

func main() {

	mCurre := make(currencyMap, 3)
		
	fmt.Println("### Конвертер валюты ###")
	
	for {

		sorcesCurr = ""
		destinationCurr = ""		
		
		mCurre["USD"] = 100
		mCurre["EUR"] = 120
		mCurre["RUB"] = 1
			
		avalCurr := availableCurrencies(&mCurre)	
		inputData(0, fmt.Sprintf("Введите исходную валюту (%s):", avalCurr), &mCurre)

		inputData(2, "Введите сумму", &mCurre)

		avalCurr = availableCurrencies(&mCurre)
		inputData(1, fmt.Sprintf("Введите валюту назначения (%s):", avalCurr), &mCurre)

		result := convert(sorcesCurr, destinationCurr, sumConv, &mCurre)

		fmt.Printf("Результат - %f %s = %f %s \n", sumConv, sorcesCurr, result, destinationCurr)

		fmt.Print("Продолжить? (y/n) ")
		ans := "" 
		fmt.Scan(&ans)
		if ans == "n" {
			break
		}
	}	
}

func inputData(mode int, messageText string, mCurr *currencyMap) {

	fmt.Println(messageText)
	switch mode {
	case 0:
		for {
			fmt.Scanln(&sorcesCurr)
			if checkInputData(sorcesCurr, *mCurr) {
				break
			}
		}			
	case 1:
		for {
			fmt.Scanln(&destinationCurr)
			if checkInputData(destinationCurr, *mCurr) {
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

func checkInputData(currency string, mCurr currencyMap) bool{

	_, found := mCurr[currency]
	if found && sorcesCurr != destinationCurr {
		return true
	}
		
	fmt.Println("Не верно указана валюта")
	avalCurr := availableCurrencies(&mCurr)
	fmt.Printf("Доступно: %s\n", avalCurr)

	return false
	
}

func availableCurrencies(mapC *currencyMap) string{
	
	avCurr := ""

	for k, _ := range *mapC{
		if k == sorcesCurr {
			continue
		}
		avCurr = avCurr + k + "/"
	}  

	return avCurr

}

func convert(firstVal string, secVal string, sum float64, calcMap *currencyMap) float64 {

	var result float64

	cm := *calcMap

	switch {
	case firstVal == "USD" && secVal == "RUB":
		result = sum * cm[firstVal]
	case firstVal == "RUB" && secVal == "USD":
		result = sum / cm[secVal]
	case firstVal == "USD" && secVal == "EUR":
		result = sum * cm[firstVal] / cm[secVal]
	case firstVal == "EUR" && secVal == "USD":
		result = sum / cm[secVal] * cm[firstVal]
	case firstVal == "EUR" && secVal == "RUB":
		result = sum * cm[firstVal]
	case firstVal == "RUB" && secVal == "EUR":
		result = sum / cm[secVal]
	default:
		result = 0.0
	}
	return result

}
