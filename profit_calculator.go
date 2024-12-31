package main

import "fmt"

func main() {

	revenue := prinText("Revenue: ")
	expenses := prinText("Expenses: ")
	taxRate := prinText("Tax Rate: ")

	ebt, profit, ratio := calculator(revenue, expenses, taxRate)

	formattedEbt := fmt.Sprintf("ebt: %.2f\n", ebt)
	formattedProfit := fmt.Sprintf("profit: %.2f\n", profit)
	formattedRatio := fmt.Sprintf("ratio: %.2f\n", ratio)

	fmt.Print(formattedEbt, formattedProfit, formattedRatio)

}

func prinText(infoText string) float64 {
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)

	return userInput
}

func calculator(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}
