package main

import (
	"fmt"
)

func main() {

	const eurToUsd float64 = 1.17
	const usdToRub float64 = 81.15
	const eurToRub = usdToRub * eurToUsd
	fmt.Println("EUR to RUB rate:", eurToRub)
	getUserInput()
	calculateCurrency(5, "sosiska", "salfetka")
}
func getUserInput() (float64, float64) {
	var value1, value2 float64
	fmt.Print("Input some: ")
	fmt.Scan(&value1)
	fmt.Print("Input some2: ")
	fmt.Scan(&value2)
	return value1, value2
}

func calculateCurrency(someInt int, currency1 string, currency2 string) float64 {
	result := 0.0
	return result
}
