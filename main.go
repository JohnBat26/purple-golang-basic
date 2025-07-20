package main

import (
	"fmt"
)

func main() {
	fmt.Println("__ Конвертер валют __")
	var usdToEur float64 = 0.85
	var usdToRub float64 = 78.5

	// usdToRub, usdToEur := getUserInput()
	eurToRub := usdToRub / usdToEur

	fmt.Println("Euro to Ruble: ", eurToRub)
}

func getUserInput() (float64, string, string) {
	var amountMoney float64
	var fromCurrency string
	var toCurrency string

	fmt.Print("Введите количество денег: ")
	fmt.Scan(&amountMoney)
	fmt.Print("Введите имя исходной валюты, например: usd, eur или rub: ")
	fmt.Scan(&fromCurrency)
	fmt.Print("Введите имя целевой валюты, например: usd, eur или rub: ")
	fmt.Scan(&toCurrency)

	return amountMoney, fromCurrency, toCurrency
}

func calculate(amount float64, fromCurrency string, toCurrency string) {

}
