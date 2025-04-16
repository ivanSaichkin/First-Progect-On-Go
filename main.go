package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	for {
		fmt.Println("калькулятор индекса массы тела")
		userKg, userHeight := GetUserInput()
		IMT, err := CalculateIMT(userKg, userHeight)
		if err != nil {
			//fmt.Println(err)
			//continue

			panic(err)
		}
		OutputResult(IMT)
		isRepeatCalculation := CheckRepeatCalculation()
		if !isRepeatCalculation {
			break
		}
	}
}

func OutputResult(IMT float64) {
	result := fmt.Sprintf("ваш индекс массы тела: %.0f", IMT)
	fmt.Println(result)
	switch {
	case IMT < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case IMT < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case IMT < 25:
		fmt.Println("У вас нормальный вес")
	case IMT < 30:
		fmt.Println("У вас избыточная масса тела")
	default:
		fmt.Println("У вас ожирение")
	}
}

func CalculateIMT(userKg float64, userHeight float64) (float64, error) {
	if userKg <= 0 || userHeight <= 0 {
		return 0, errors.New("не указан весь или рост")
	}
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT, nil
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

func CheckRepeatCalculation() bool {
	var userChoice string

	fmt.Print("Хотите сделать еще расчет (y/n): ")
	fmt.Scan(&userChoice)

	if userChoice == "y" || userChoice == "Y" {
		return true
	}
	return false
}
