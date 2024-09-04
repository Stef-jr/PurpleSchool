package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	
	fmt.Println("--- Калькулятор функций ---")
	oper := scanOperation()
  fmt.Printf("Используется операция %s\n", oper)
	intSlice := scanDatas()
	resault := calculateValues(oper, intSlice)
	fmt.Printf("Результат операции %s = %.2f\n", oper, resault)

}

func scanOperation() string {
	
	operation := ""

	fmt.Print("Введите операцию: (AVG/SUM/MED) ")
	fmt.Scan(&operation)
	return strings.TrimSpace(operation)

}

func scanDatas() []float64{
	
	fmt.Println("Введите набор числе разделяя их ','.")
	reader := bufio.NewReader(os.Stdin)
	inputstr, _ := reader.ReadString('\n')
	inputstr = strings.TrimSpace(inputstr)
	
	trimstring := strings.Split(strings.ReplaceAll(inputstr, " ", ""), ",")
	
	massData := make([]float64, len(trimstring))
	
	for index, s := range trimstring {
		num, err := strconv.ParseFloat(s, 64)
		if err != nil {
			fmt.Println("Ошибка преобразования: ", err)
			continue
		}
		massData[index] = num
	}
	
	return massData

}

func calculateValues(oper string, numSlice []float64) float64 {
	
	switch oper {
	case "AVG":
		sum := 0.0
		for _, value := range numSlice {
			sum += value 
		}		
		return sum / float64(len(numSlice))
	case "SUM":
		sum := 0.0
		for _, value := range numSlice {
			sum += value
		}
		return sum
	case "MED":
		sort.Float64s(numSlice)
		n := len(numSlice)
		if n%2 == 0 {
			return (numSlice[n/2-1] + numSlice[n/2]) / 2
		} else {
			return numSlice[n/2]
		}		
	default:
		fmt.Println("Нет такой операции!")
		return 0.0
	}

}
