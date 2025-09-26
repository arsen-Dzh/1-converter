package main

import (
	"fmt"
)

const usd = "usd"
const eur = "eur"
const rub = "rub"

type coursesCurrencyMap = map[string]map[string]float64

func main() {
	coursesCurrency := coursesCurrencyMap{
		"rub": {
			"usd": 0.01,
			"eur": 0.01,
		},
		"usd": {
			"rub": 83.99,
			"eur": 0.85,
		},
		"eur": {
			"rub": 98.6,
			"usd": 1.17,
		},
	}

	userCurrency, targetCurrency, userSum := getUserValue()
	resultSum := sumCount(userCurrency, targetCurrency, userSum, &coursesCurrency)
	fmt.Print("Вы получите: ", resultSum)
}

func getUserValue() (string, string, float64) {
	var userCurrency string
	var targetCurrency string
	var userSum float64

	for {
		fmt.Println("Какую валюту вы хотите обменять usd, eur, rub")
		fmt.Scan(&userCurrency)
		if userCurrency == usd || userCurrency == eur || userCurrency == rub {
			break
		}
		fmt.Println("Выберите валюту из списка")
		continue
	}

	for {
		fmt.Println("Какую сумму вы хотите обменять?")
		fmt.Scan(&userSum)
		if userSum > 0 {
			break
		}
		fmt.Println("Введите число больше нуля")
		continue
	}

	for {
		fmt.Println("Какую валюту вы хотите получить usd eur rub?")
		fmt.Scan(&targetCurrency)
		if targetCurrency == userCurrency {
			fmt.Println("Ваша валюта должна отличаться от целевой")
			continue
		} else if targetCurrency == usd || targetCurrency == eur || targetCurrency == rub && targetCurrency != userCurrency {
			break
		} else {
			fmt.Println("Выберите валюту из списка")
			continue
		}
	}
	return userCurrency, targetCurrency, userSum
}

func sumCount(userCurrency string, targetCurrency string, userSum float64, coursesCurrency *coursesCurrencyMap) float64 {
	var result float64

	shouldBreak := false
	for currency, courses := range *coursesCurrency {
		if shouldBreak {
			break
		}

		if currency == userCurrency {
			for key, val := range courses {
				if key == targetCurrency {
					result = userSum * val
					shouldBreak = true
					break
				}
			}
		}
	}

	return result
}
