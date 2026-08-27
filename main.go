package main

import "fmt"

const (
	USDInEUR float64 = 0.86
	USDInRUB float64 = 83.65
	EURInRUB float64 = USDInRUB / USDInEUR
)

func main() {
	strUser := writeUser()
	print(strUser)
}

func writeUser() string {
	var str string
	fmt.Scan(&str)
	return str
}

// func convertersValue(num, val1, val2 float64) string {

// }