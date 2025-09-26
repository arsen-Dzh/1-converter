package main

import (
	"fmt"
)

const usd = "usd"
const eur = "eur"
const rub = "rub"

func main() {
	userCurrency, targetCurrency, userSum := getUserValue()
	resultSum := sumCount(userCurrency, targetCurrency, userSum)
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

func sumCount(userCurrency string, targetCurrency string, userSum float64) float64 {
	var result float64
	switch {
	case userCurrency == rub && targetCurrency == usd:
		result = userSum * 0.01
	case userCurrency == rub && targetCurrency == eur:
		result = userSum * 0.01
	case userCurrency == usd && targetCurrency == rub:
		result = userSum * 83.99
	case userCurrency == usd && targetCurrency == eur:
		result = userSum * 0.85
	case userCurrency == eur && targetCurrency == rub:
		result = userSum * 98.6
	case userCurrency == eur && targetCurrency == usd:
		result = userSum * 1.17
	}

	return result
}
