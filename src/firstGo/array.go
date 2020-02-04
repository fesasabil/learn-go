package main

import "fmt"

func main() {
	var names [4] string

	names[0] = "Trafalgar"
	names[1] = "D"
	names[2] = "Water"
	names[3] = "Law"

	fmt.Println(names[0], names[1], names[2], names[3])

	//  Inisialisasi Nilai Awal Array
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	// Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen
	var numbers = [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t:", numbers)
	fmt.Println("Jumlah elament \t:", len(numbers))

	// Array Multidimensi
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 4, 5}, {5, 6, 7}}

	fmt.Println("numbers", numbers1)
	fmt.Println("numbers", numbers2)

	// Perulangan Elemen Array Menggunakan Keyword for
	var fruits1 = [4]string{"apple", "grape", "banana", "melon"}

	for i := 0; i < len(fruits1); i++ {
		fmt.Printf("element %d : %s\n", i, fruits1[i])
	}

	// Perulangan Elemen Array Menggunakan Keyword for - range
	var fruits2 = [4]string{"apple", "grape", "banana", "melon"}

	for i, fruits := range fruits2 {
		fmt.Printf("element %d : %s\n", i, fruits)
	}

	//  Penggunaan Variabel Underscore _ Dalam for - range
	var fruits3 = [4]string{"apple", "grape", "banana", "melon"}

	for _, fruit := range fruits3 {
		fmt.Printf("nama buah : %s\n", fruit)
	}

	// Alokasi Elemen Array Menggunakan Keyword make
	var food = make([]string, 2)
	food[0] = "Pizza"
	food[1] = "Humberger"

	fmt.Println(food)
}