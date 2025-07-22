package main

import (
	"fmt"
	"strings"
)

const UsDollar = "usd"
const Euro = "eur"
const RussianRuble = "rub"

func main() {
	fmt.Println("__ Конвертер валют __")

	var usdToEur = 0.85
	var usdToRub = 78.5
	var eurToRub = usdToRub / usdToEur

	for {
		amountMoney, fromCurrency, toCurrency := getUserInput()

		var targetAmountMoney float64

		switch {
		case fromCurrency == UsDollar && toCurrency == Euro:
			targetAmountMoney = amountMoney * usdToEur
		case fromCurrency == UsDollar && toCurrency == RussianRuble:
			targetAmountMoney = amountMoney * usdToRub
		case fromCurrency == Euro && toCurrency == UsDollar:
			targetAmountMoney = amountMoney / usdToEur
		case fromCurrency == RussianRuble && toCurrency == UsDollar:
			targetAmountMoney = amountMoney / usdToRub
		case fromCurrency == Euro && toCurrency == RussianRuble:
			targetAmountMoney = amountMoney * eurToRub
		case fromCurrency == RussianRuble && toCurrency == Euro:
			targetAmountMoney = amountMoney / eurToRub
		default:
			fmt.Println("Незнакомая комбинация валют")
		}

		fmt.Println("Результаты конвертации:")
		result := fmt.Sprintf("%0.2f %s => %.2f %s", amountMoney, fromCurrency, targetAmountMoney, toCurrency)
		fmt.Println(result)

		fmt.Println("Продолжить работу? y/n")
		var answer string
		fmt.Scan(&answer)

		if strings.EqualFold(answer, "y") {
			continue
		} else {
			break
		}
	}
}

func getUserInput() (float64, string, string) {
	var amountMoney float64
	var fromCurrency string
	var toCurrency string

	for {
		message := fmt.Sprintf("Введите имя исходной валюты, например: %s, %s или %s: ", UsDollar, Euro, RussianRuble)
		fmt.Println(message)
		fmt.Scan(&fromCurrency)
		if fromCurrency != UsDollar && fromCurrency != Euro && fromCurrency != RussianRuble {
			message = fmt.Sprintf("Вы должны ввести одно из трех значений: %s, %s или %s, Попробуйте снова.", UsDollar, Euro, RussianRuble)
			fmt.Println(message)
		} else {
			break
		}
	}

	for {
		fmt.Print("Введите сумму средств для конвертации: ")
		_, err := fmt.Scan(&amountMoney)
		if err != nil {
			fmt.Println("Некорректное число! ", err)
			continue
		}
		break
	}

	for {
		// Формируем приглашение с примерами допустимых валют
		prompt := fmt.Sprintf("Введите имя целевой валюты, например: %s, %s или %s: ",
			UsDollar, Euro, RussianRuble)
		fmt.Print(prompt) // Используем Print вместо Println для ввода на той же строке

		// Читаем ввод пользователя
		fmt.Scan(&toCurrency)

		// Проверяем корректность ввода
		switch {
		case toCurrency != UsDollar && toCurrency != Euro && toCurrency != RussianRuble:
			fmt.Printf("Вы должны ввести одно из трех значений: %s, %s или %s. Попробуйте снова.\n",
				UsDollar, Euro, RussianRuble)
		case toCurrency == fromCurrency:
			fmt.Println("Ошибка: целевая валюта совпадает с исходной. Повторите попытку.")
		default:
			return amountMoney, fromCurrency, toCurrency
		}
	}
}
