package main

import (
	"fmt"
)

func main() {
	fmt.Println("__ Конвертер валют __")

	usdToRub, usdToEur := getUserInput()
	eurToRub := usdToRub / usdToEur

	fmt.Println("Euro to Ruble: ", eurToRub)
}

func getUserInput() (float64, float64) {
	var usdToEur float64 = 0.85
	var usdToRub float64 = 78.5

	fmt.Print("Введите коэффициент преобразования доллара в евро: ")
	fmt.Scan(&usdToEur)
	fmt.Print("Введите коэффициент преобразования долларав рубль: ")
	fmt.Scan(&usdToRub)

	return usdToEur, usdToRub
}

func calculate(amount float64, fromCurrency float64, toCurrency float64) {

}
