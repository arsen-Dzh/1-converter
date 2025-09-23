package main

import "fmt"

func main() {
	const usdToEur float64 = 0.84
	const usdToRub float64 = 83.35
	const eurToRub = usdToRub / usdToEur

	fmt.Print(eurToRub)
}
