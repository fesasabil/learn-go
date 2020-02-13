package main

import "fmt"

/** 
	Definisi Closure adalah sebuah fungsi yang bisa disimpan dalam variabel. Dengan menerapkan konsep tersebut, kita bisa membuat fungsi didalam fungsi, atau bahkan membuat fungsi yang mengembalikan fungsi.

	Closure merupakan anonymous function atau fungsi tanpa nama. Biasa dimanfaatkan untuk membungkus suatu proses yang hanya dipakai sekali atau dipakai pada blok tertentu saja.

	// Penggunaan Template String %v
		Template %v digunakan untuk menampilkan segala jenis data. Bisa array, int, float, bool, dan lainnya.

	//  Closure Sebagai Nilai Kembalian
		Salah satu keunikan closure lainnya adalah bisa dijadikan sebagai nilai balik fungsi, cukup aneh memang, tapi pada suatu kondisi teknik ini sangat membantu. Di bawah ini disediakan sebuah fungsi bernama findMax(), fungsi ini salah satu nilai kembaliannya berupa closure.

	Sedikit tentang fungsi findMax(), fungsi ini digunakan untuk mencari banyaknya angka-angka yang nilainya di bawah atau sama dengan angka tertentu. Nilai kembalian pertama adalah jumlah angkanya. Nilai kembalian kedua berupa closure yang mengembalikan angka-angka yang dicari. Berikut merupakan contoh implementasi fungsi tersebut.
*/

// func main()  {
// 	var getMinMax = func (n []int) (int,int)  {
// 		var min, max int

// 		for i, e := range n {
// 			switch  {
// 			case i == 0:
// 				max, min = e, e
// 			case e > max:
// 				max = e
// 			case e < min:
// 				min = e
// 			}
// 		}

// 		return min, max
// 	}

// 	var numbers = []int{2, 3, 4, 3, 4, 2, 4, 3}
// 	var min, max = getMinMax(numbers)
// 	fmt.Printf("data : %v\nmin	: %v\nmax	: %v\n", numbers, min, max)

// }

// Immediately-Invoked Function Expression (IIFE)
// func main() {
// 	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

// 	var newNumbers = func(min int) []int { // IIFE
// 		var r []int
// 		for _, e := range numbers {
// 			if e <  min {
// 				continue
// 			}

// 			r = append(r, e)
// 		}

// 		return r
// 	} (3)

// 	fmt.Println("original number :", numbers)
// 	fmt.Println("filtered number :", newNumbers)
// }

//  Closure Sebagai Nilai Kembalian
func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func()[]int {
		return res
	}
}

func main()  {
	var max = 3
	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var howMany,  getNumbers = findMax(numbers, max)
	var theNumbers = getNumbers()

	fmt.Println("numbers\t:", numbers)
	fmt.Printf("find \t: %d\n\n", max)

	fmt.Println("found \t:", howMany)	// 9
	fmt.Println("value \t:", theNumbers)	// [2 3 0 3 2 0 2 0 3]
}