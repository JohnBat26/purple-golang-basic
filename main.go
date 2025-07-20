package main

import (
	"fmt"
)

func main() {
	const usdToEur = 0.85
	const usdToRub = 78.5
	const eurToRub = usdToRub / usdToEur

	fmt.Println("Euro to Ruble: ", eurToRub)
}
