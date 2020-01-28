package main

import "fmt"

func main() {
	// Tipe Data Numerik Non-Desimal
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644
	// Tipe data numerik decimal
	var decimalNumber = 2.62
	// Tipe data Bool(Boolean)
	var exist bool = true
	// Tipe data string
	var message string = "Hallo"
	// Tipe data string menggunakan backticks
	var pesan = `Nama saya "John Wick".
	Salam Kenal.
	Mari belajar "Golang".`

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negative: %d\n", negativeNumber)
	fmt.Printf("bilangan decimal: %f\n", decimalNumber)
	fmt.Printf("bilangan decimal: %.3f\n", decimalNumber)
	fmt.Printf("exist? %t \n", exist)
	fmt.Printf("message: %s \n", message)
	fmt.Printf(pesan)
}