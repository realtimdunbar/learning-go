package main

import "fmt"

func main() {
	fmt.Println("Enter your temperature in Farenheit: ")
	var farenheit float64
	fmt.Scanf("%f", &farenheit)

	celsius := (farenheit - 32) * 5 / 9

	fmt.Println(celsius)
}
