package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("__ Калькулятор по списку __")
	operation := scanOperation()
	numbers := scanNumbers()
	calculate(operation, numbers)
}

func calculate(operation string, numbers []float64) {
	switch operation {
	case "avg":
		fmt.Println("Среднее значение элементов равно: ", calculateAverage(numbers))
	case "sum":
		fmt.Println("Сумма элементов равна: ", calculateSum(numbers))
	case "med":
		fmt.Println("Медиана элементов равна: ", calculateMedian(numbers))
	}
}

func scanOperation() string {
	fmt.Println("Введите одну из операций AVG -> среднее, SUM -> сумма, MED -> медиана:")
	var operation string
	for {
		_, err := fmt.Scan(&operation)
		if err != nil {
			fmt.Println("Ошибка при вводе данных:", err)
			panic(err)
		}

		operation = strings.ToLower(operation)

		if operation != "avg" && operation != "sum" && operation != "med" {
			fmt.Println("Введена некорректная операция, попробуйте ещё раз")
			continue
		}
		break
	}
	return operation
}

func scanNumbers() []float64 {
	fmt.Println("Введите список чисел разделенных запятыми:")

	var numbersFromUser string
	fmt.Scan(&numbersFromUser)

	numbersAsString := strings.Split(numbersFromUser, ",")
	numbersAsFloats := make([]float64, len(numbersAsString))

	for i, s := range numbersAsString {
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			fmt.Printf("Ошибка преобразования строки '%s': %v\n", s, err)
			continue
		}

		numbersAsFloats[i] = f
	}

	return numbersAsFloats
}

func calculateSum(numbers []float64) float64 {
	sum := 0.0
	for _, f := range numbers {
		sum += f
	}
	return sum
}

func calculateAverage(numbers []float64) float64 {
	sum := 0.0
	for _, f := range numbers {
		sum += f
	}
	avg := sum / float64(len(numbers))
	return avg
}

func calculateMedian(numbers []float64) float64 {
	sort.Float64s(numbers)

	length := len(numbers)
	var median float64

	if length == 0 {
		return 0
	}

	if length%2 == 1 {
		median = numbers[length/2]
	} else {
		median = (numbers[length/2-1] + numbers[length/2]) / 2
	}

	return median
}
