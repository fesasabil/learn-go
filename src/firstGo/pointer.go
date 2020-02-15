package main

import "fmt"

/** 
	Pointer adalah reference atau alamat memori.Variabel pointer berarti variabel yang berisi alamat memori suatu nilai.Sebagai contoh sebuah variabel bertipe integer memiliki nilai 4, maka yang dimaksud pointer adalah alamat memori dimana nilai 4 disimpan, bukan nilai 4 itu sendiri.

	Variabel-variabel yang memiliki reference atau alamat pointer yang sama, saling berhubung satu sama lain dan nilainya pasti sama. Ketika ada perubahan nilai, maka akan memberikan efek kepada variabel lain(yang referensi-nya sama) yaitu nilainya ikut berubah.

	Variabel bertipe pointer ditandai dengan adanya tanda asterisk (*) tepat sebelum penulisan tipe data ketika deklarasi.

	Nilai default variabel pointer adalah nil (kosong). Variabel pointer tidak bisa menampung nilai yang bukan pointer, dan sebaliknya variabel biasa tidak bisa menampung nilai pointer.

	Ada dua hal penting yang perlu diketahui mengenai pointer:

	1. Variabel biasa bisa diambil nilai pointernya, caranya dengan menambahkan tanda ampersand (&) tepat sebelum nama variabel. Metode ini disebut dengan referencing.
	2. Dan sebaliknya, nilai asli variabel pointer juga bisa diambil, dengan cara menambahkan tanda asterisk (*) tepat sebelum nama variabel. Metode ini disebut dengan dereferencing.
*/

// func  main()  {
// 	var numberA int = 4
// 	var numberB *int = &numberA

// 	fmt.Println("numberA (value) :", numberA)	// 4
// 	fmt.Println("numberA (address) :", &numberA)	// 0xc0000160b8

// 	fmt.Println("numberB (value) :", *numberB)	// 4
// 	fmt.Println("numberB (address) :", numberB)	// 0xc0000160b8

// 	fmt.Println("")

// 	numberA = 5

// 	fmt.Println("numberA (value) :", numberA)
// 	fmt.Println("numberA (address) :", &numberA)
// 	fmt.Println("numberB (value) :", *numberB)
// 	fmt.Println("numberB (address) :", numberB)
// }

// Parameter Pointer

func main() {
	var number = 4
	fmt.Println("before :", number) // 4

	change(&number, 10)
	fmt.Println("after :", number) // 10
}

func change(original *int, value int)  {
	*original = value
}