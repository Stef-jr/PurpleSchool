package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var menu = map[string]func([]float64) float64{
	"1": average,
	"2": sum,
	"3": mediana,
}

var variantMenu = []string{
	"1. Среднее",
	"2. Сумма",
	"3. Медиана",
	"4. Выход",
	"Выберите операцию: ",
}

func main() {

	fmt.Println("--- Калькулятор функций ---")
Menu:
	for {
		oper := scanOperation(variantMenu...)
		menuFunc := menu[oper]
		if menuFunc == nil {
			break Menu
		}
		intSlice := scanData()
		res := menuFunc(intSlice)
		fmt.Printf("Результат = %.2f\n", res)
	}
}

func scanOperation(operations ...string) string {

	operation := ""
	for i, line := range operations {
		if i == len(operations)-1 {
			fmt.Printf("%v", line)
		} else {
			fmt.Println(line)
		}
	}

	fmt.Scan(&operation)
	return strings.TrimSpace(operation)

}

func scanData() []float64 {

	fmt.Println("Введите набор числе разделяя их ','.")
	reader := bufio.NewReader(os.Stdin)
	inputStr, _ := reader.ReadString('\n')
	inputStr = strings.TrimSpace(inputStr)

	trimStr := strings.Split(strings.ReplaceAll(inputStr, " ", ""), ",")

	massData := make([]float64, len(trimStr))

	for index, s := range trimStr {
		num, err := strconv.ParseFloat(s, 64)
		if err != nil {
			fmt.Println("Ошибка преобразования: ", err)
			continue
		}
		massData[index] = num
	}

	return massData

}

func average(numSlice []float64) float64 {
	sum := 0.0
	for _, value := range numSlice {
		sum += value
	}
	return sum / float64(len(numSlice))
}

func sum(numSlice []float64) float64 {
	sum := 0.0
	for _, value := range numSlice {
		sum += value
	}
	return sum
}

func mediana(numSlice []float64) float64 {
	sort.Float64s(numSlice)
	n := len(numSlice)
	if n%2 == 0 {
		return (numSlice[n/2-1] + numSlice[n/2]) / 2
	} else {
		return numSlice[n/2]
	}
}
