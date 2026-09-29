package main

import "fmt"

func main() {
	var fahrenheit, celsius float64

	fmt.Print("Masukkan suhu Fahrenheit: ")
	fmt.Scan(&fahrenheit)

	celsius = (fahrenheit - 32) * 5.0 / 9.0

	fmt.Printf("Suhu Celsius = %.0f\n", celsius)
}