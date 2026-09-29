package main

import "fmt"

func main() {
	var r, luas float64

	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&r)

	luas = (22.0 / 7.0) * r * r

	fmt.Printf("Luas lingkaran = %.1f\n", luas)
}