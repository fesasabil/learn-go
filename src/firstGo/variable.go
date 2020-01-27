package main

import "fmt"

func main() {
	// Deklarasi multi variabel
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	// variable tanpa menuliskan var
	address := "Bogor"
	// mengubah variable address
	address = "Khazakstan"
	fmt.Printf("Hello %s %s %s %s!\n", first, second, third, address)
}