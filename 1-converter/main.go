package main

import (
	"fmt"
)

func main() {
	const eurToUsd float64 = 1.17
	const usdToRub float64 = 81.15
	const eurToRub = usdToRub * eurToUsd
	fmt.Println("EUR to RUB rate:", eurToRub)
}
