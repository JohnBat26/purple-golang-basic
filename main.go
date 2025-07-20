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

func getUserInput() (float64, float64, float64) {
	var amountMoney float64
	var fromCurrency float64
	var toCurrency float64

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
