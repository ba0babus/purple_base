package main

import (
	"fmt"
)

func main() {
	rates := map[string]float64{
		"USD_RUB": 81.15,
		"USD_EUR": 0.85,
		"RUB_USD": 0.012,
		"RUB_EUR": 0.011,
		"EUR_RUB": 92.12,
		"EUR_USD": 1.17,
	}
	fromCurrency, amount, toCurrency := getUserInput()
	calculateCurrency(fromCurrency, amount, toCurrency, &rates)
}

func getCurrency(currencyType string, currencyMap *map[string]string) string {
	currency := ""
	for { // выбираем исходную валюту
		fmt.Printf("Choose the currency number %v convert (e.g.: 2): \n", currencyType)
		fmt.Printf("USD = 1\nEUR = 2\nRUB = 3\n")
		fmt.Scan(&currency)
		_, exists := (*currencyMap)[currency]
		if exists {
			return currency
		} else {
			fmt.Printf("Please, type a number from 1 to 3. Try again...\n\n\n")
		}
	}
}

func getUserInput() (string, float64, string) {
	var amount float64
	currencyMap := map[string]string{
		"1": "USD",
		"2": "EUR",
		"3": "RUB",
	}

	fromCurrency := currencyMap[getCurrency("from", &currencyMap)]

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

	toCurrency := currencyMap[getCurrency("to", &currencyMap)]

	return fromCurrency, amount, toCurrency
}

func calculateCurrency(fromCurrency string, amount float64, toCurrency string, rates *map[string]float64) {

	var result float64

	// Если валюты одинаковые
	if fromCurrency == toCurrency {
		result = amount
		fmt.Printf("Total: %0.2f\n", result)
		return
	}

	// Формируем ключ для поиска в map
	key := fromCurrency + "_" + toCurrency

	// Ищем курс в map
	if rate, exists := (*rates)[key]; exists {
		result = amount * rate
		fmt.Printf("Total: %0.2f\n", result)
	} else {
		fmt.Println("Invalid currency pair")
	}
}
