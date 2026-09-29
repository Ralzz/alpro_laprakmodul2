package main

import "fmt"

func main() {
	var nama, nim, kelas string

	// Meminta masukan satu per satu dengan teks petunjuk
	fmt.Print("Masukkan Nama  : ")
	fmt.Scanln(&nama)

	fmt.Print("Masukkan NIM   : ")
	fmt.Scanln(&nim)

	fmt.Print("Masukkan Kelas : ")
	fmt.Scanln(&kelas)

	// Menampilkan resume singkat mahasiswa
	fmt.Printf("\nPerkenalkan saya adalah %s, salah satu mahasiswa Prodi PS1IF-14 dari kelas %s dengan NIM %s.\n", nama, kelas, nim)
}