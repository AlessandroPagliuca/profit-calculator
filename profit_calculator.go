package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	revenue, err := getUserInput("Revenue: ")

	if err != nil {
		fmt.Println(err)
		//panic(err)
		return
	}

	expenses, err := getUserInput("Expenses: ")

	if err != nil {
		fmt.Println(err)
		//panic(err)
		return
	}

	taxRate, err := getUserInput("Tax Rate: ")

	if err != nil {
		fmt.Println(err)
		//panic(err)
		return
	}

	ebt, profit, ratio := calculator(revenue, expenses, taxRate)

	fmt.Printf("%.2f\n", ebt)
	fmt.Printf("%.2f\n", profit)
	fmt.Printf("%.2f\n", ratio)

	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	numText := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(numText), 0644)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("value not valid")
	}

	return userInput, nil
}

func calculator(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return
}
