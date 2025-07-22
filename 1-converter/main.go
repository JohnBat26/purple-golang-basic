package main

import (
	"fmt"
	"strings"
)

const (
	UsDollar     = "usd"
	Euro         = "eur"
	RussianRuble = "rub"
)

type exchangeRates map[string]float64

func main() {
	currencies := exchangeRates{
		"usd-eur": 0.85,
		"eur-usd": 1 / 0.85,
		"usd-rub": 78.5,
		"rub-usd": 1 / 78.5,
		"eur-rub": 92.3,
		"rub-eur": 1 / 92.3,
	}

	fmt.Println("__ Конвертер валют __")

	for {
		amountMoney, fromCurrency, toCurrency := getUserInput()

		var targetAmountMoney float64

		key := fromCurrency + "-" + toCurrency
		if rate, exists := currencies[key]; exists {
			targetAmountMoney = amountMoney * rate
		}

		fmt.Println("Результаты конвертации:")
		result := fmt.Sprintf("%0.2f %s => %.2f %s", amountMoney, fromCurrency, targetAmountMoney, toCurrency)
		fmt.Println(result)

		fmt.Println("Продолжить работу? y/n")

		var answer string

		fmt.Scan(&answer)

		if strings.EqualFold(answer, "y") {
			continue
		}

		break
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
