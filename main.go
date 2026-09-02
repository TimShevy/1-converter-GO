package main

import (
	"fmt"
)

const (
	usdInRub float64 = 83.65
	usdInEur float64 = 0.86
	euRInRub float64 = usdInRub / usdInEur
)

var currencyRates = &map[string]float64{
	"USD": 1,
	"EUR": 1 / usdInEur,
	"RUB": 1 / usdInRub,
}
func main() {
	
	fmt.Println("___ Конвертер валюты ___")
	for {
		userCurrency := getCurrency()
		userAmount := getAmount()
		convertCurrency := getConvertCurrency(userCurrency)
		
		finish := finishConvert(userCurrency, convertCurrency, userAmount)
		fmt.Printf("Результат: %.2f\n", finish)
		
		
		isRepeateConvertation := checkRepeatConvertation()
		if !isRepeateConvertation {
			fmt.Print("Всего доброго, досвидания!")
			break
		}
	}
}

func finishConvert(userCurrency, convertCurrency string, userAmount float64) float64 {
	return userAmount * (*currencyRates)[userCurrency] / (*currencyRates)[convertCurrency]
}
// func finishConvert(userCurrency, convertCurrency string, userAmount float64) float64 {
// 	var result float64
// 	switch {
// 	case userCurrency == "USD" && convertCurrency == "RUB":
// 		result = userAmount * usdInRub
// 	case userCurrency == "USD" && convertCurrency == "EUR":
// 		result = userAmount * usdInEur
// 	case userCurrency == "RUB" && convertCurrency == "USD":
// 		result = userAmount / usdInRub
// 	case userCurrency == "RUB" && convertCurrency == "EUR":
// 		result = userAmount / euRInRub
// 	case userCurrency == "EUR" && convertCurrency == "RUB":
// 		result = userAmount * euRInRub
// 	case userCurrency == "EUR" && convertCurrency == "USD":
// 		result = userAmount / usdInEur
// 	default: fmt.Println("Что то пошло не так!")
// 		result = 0
// 	}
// 	return result
// }

func getConvertCurrency(userCurrency string) string {
	var convertCurrency string
	for {
		fmt.Printf("Введите целевую валюту, кроме '%s': ", userCurrency)
		_, err := fmt.Scan(&convertCurrency)
		if err != nil || convertCurrency == userCurrency{
			fmt.Println("Ошибка ввода. Введите другую валюту.")
			continue
		}
		if convertCurrency == "RUB" || convertCurrency == "USD" || convertCurrency == "EUR" {
			return convertCurrency
		}
		fmt.Println("Такой валюты нет. Введите валюту ещё раз.")
	}
}

func getCurrency() string {
	var currency string
	for {
		fmt.Print("Введите исходную валюту. 'USD', 'EUR' или 'RUB': ")
		_, err := fmt.Scan(&currency)
		if err != nil {
			fmt.Println("Ошибка ввода.")
			continue
		}
		if currency == "RUB" || currency == "USD" || currency == "EUR" {
			return currency
		}
		fmt.Println("Ошибка ввода. Попробуйте ещё раз.")
	}
}

func getAmount() float64 {
	var amount float64
	for {
		fmt.Print("Введите сумму: ")
		_, err := fmt.Scan(&amount)
		if err != nil {
			fmt.Println("Ошибка ввода. Введите число!")
			continue
		}
		if amount > 0 {
			return amount
		}
		fmt.Println("Ошибка ввода. Нужно ввести число больше ноля.")
	}
}

func checkRepeatConvertation() bool {
	var userChoice string
	fmt.Print("Вы хотите сделать еще конвертацию? (y/n): ")
	fmt.Scan(&userChoice)
	if userChoice == "Y" || userChoice == "y" {
		return true
	}
	return false
}