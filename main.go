package main

import "fmt"

func main() {
	const usdToEur float64 = 0.84
	const usdToRub float64 = 83.35
	const eurToRub = usdToRub / usdToEur

	fmt.Println(eurToRub)
	getUserValue()
}

func getUserValue() string {
	var userValue string

	fmt.Print("Введите текст: ")
	fmt.Scan(&userValue)

	return userValue
}

func fooCount(numFoo, usd, rub float64) float64 {
	return 1
}
