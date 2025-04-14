package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("калькулятор индекса массы тела")
	userKg, userHeight := GetUserInput()
	IMT := CalculateIMT(userKg, userHeight)
	OutputResult(IMT)
}

func OutputResult(IMT float64) {
	result := fmt.Sprintf("ваш индекс массы тела: %.0f", IMT)
	fmt.Print(result)
}

func CalculateIMT(userKg float64, userHeight float64) float64 {
	const IMTPower = 2
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT
}

func GetUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64

	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)

	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)

	return userKg, userHeight
}
