package main

import (
	"fmt"
)

func main() {
	calculateCurrency(getUserInput())
}

func getUserInput() (string, float64, string) {
	var fromCurrency, toCurrency string
	var amount float64
fromCurrencyLoop:
	for { // выбираем исходную валюту
		fmt.Printf("Choose the currency number from convert (e.g.: 2): \n")
		fmt.Printf("USD = 1\nEUR = 2\nRUB = 3\n")
		fmt.Scan(&fromCurrency)
		switch fromCurrency {
		case "1":
			fromCurrency = "USD"
			break fromCurrencyLoop
		case "2":
			fromCurrency = "EUR"
			break fromCurrencyLoop
		case "3":
			fromCurrency = "RUB"
			break fromCurrencyLoop
		default:
			fmt.Printf("Please, type a number from 1 to 3. Try again...\n\n\n")
			continue fromCurrencyLoop
		}
	}

amountLoop:
	for { // выбираем количество
		fmt.Printf("Enter the amout of currency to convert: ")
		fmt.Scan(&amount)
		if amount < 0 {
			fmt.Printf("The amount of currency cannot be less than 0. Try again...\n\n\n")
			continue amountLoop
		} else {
			break amountLoop
		}
	}
toCurrencyLoop:
	for { // выбираем исходную валюту
		fmt.Printf("Choose the currency number to convert (e.g.: 2): \n")
		fmt.Printf("USD = 1\nEUR = 2\nRUB = 3\n")
		fmt.Scan(&toCurrency)
		switch toCurrency {
		case "1":
			toCurrency = "USD"
			break toCurrencyLoop
		case "2":
			toCurrency = "EUR"
			break toCurrencyLoop
		case "3":
			toCurrency = "RUB"
			break toCurrencyLoop
		default:
			fmt.Printf("Please, type a number from 1 to 3. Try again...\n\n\n")
			continue toCurrencyLoop
		}
	}
	return fromCurrency, amount, toCurrency
}

func calculateCurrency(fromCurrency string, amount float64, toCurrency string) {
	const eurToUsd = 1.17
	const eurToRub = 94.89
	const rubToEur = 0.011
	const rubToUsd = 0.012
	const usdToEur = 0.85
	const usdToRub = 81.15
	var result float64

	if fromCurrency == "USD" && toCurrency == "RUB" {
		result = amount * usdToRub
	} else if fromCurrency == "USD" && toCurrency == "EUR" {
		result = amount * usdToEur
	} else if fromCurrency == "RUB" && toCurrency == "USD" {
		result = amount * rubToUsd
	} else if fromCurrency == "RUB" && toCurrency == "EUR" {
		result = amount * rubToEur
	} else if fromCurrency == "EUR" && toCurrency == "RUB" {
		result = amount * eurToRub
	} else if fromCurrency == "EUR" && toCurrency == "USD" {
		result = amount * eurToUsd
	} else if fromCurrency == toCurrency {
		result = amount
	} else {
		fmt.Println("Invalid currency pair")
		return
	}

	fmt.Printf("Total: %0.2f\n", result)
}
