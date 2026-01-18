package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var menu = map[string]func(){
	"1": getAVG,
	"2": getSUM,
	"3": getMED,
}

var menuVariants = []string{
	"1. Найти средрее",
	"2. Найти сумму",
	"3. Найти медиану",
	"4. Выход",
	"Выберите вариант:",
}

func main() {
	getUserOperation()
}

func getUserOperation() {
	var operation string
Menu:
	for {
		for i := range menuVariants {
			fmt.Println(menuVariants[i])
		}
		fmt.Scan(&operation)
		mathFuntion := menu[operation]
		if mathFuntion == nil {
			break Menu
		}
		mathFuntion()
	}
}

func getUserInput() []int {
	var userInput string
	var result []int
	fmt.Printf("Enter some comma separated values: ")
	fmt.Scan(&userInput)
	separatedString := strings.Split(userInput, ",")
	for _, value := range separatedString {
		num, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("ERROR: %v is not a number\n", value)
			continue
		}
		result = append(result, num)
	}
	return result
}

func getAVG() {
	userInput := getUserInput()
	var result int
	length := len(userInput)
	for _, v := range userInput {
		result += v
	}
	result = result / length
	fmt.Printf("Average is: %v\n\n\n", result)
}

func getMED() {
	userInput := getUserInput()
	var result int
	sort.Ints(userInput[:])
	if len(userInput)%2 == 0 {
		medianRightIndex := (len(userInput) / 2)
		medianLeftIndex := (len(userInput) / 2) - 1
		result = (userInput[medianRightIndex] + userInput[medianLeftIndex]) / 2
	} else {
		medianIndex := (len(userInput) / 2)
		result = userInput[medianIndex]
	}
	fmt.Printf("Median is: %v\n\n\n", result)
}

func getSUM() {
	userInput := getUserInput()
	var result int
	for _, v := range userInput {
		result += v
	}
	fmt.Printf("Summ is: %v\n\n\n", result)
}
