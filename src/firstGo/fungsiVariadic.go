/** 
	variadic function atau pembuatan fungsi dengan parameter sejenis yang tak terbatas.Maksud tak terbatas disini adalah jumlah parameter yang disisipkan ketika pemanggilan fungsi bisa berapa saja.

	Parameter variadic memiliki sifat yang mirip dengan slice. Nilai dari parameter-parameter yang disisipkan bertipe data sama, dan ditampung oleh sebuah variabel saja. Cara pengaksesan tiap datanya juga sama, dengan menggunakan index.
*/

package main

// import "fmt"

// func main() {
// 	var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
// 	var msg = fmt.Sprintf("Rata - rata : %.2f", avg)
// 	fmt.Println(msg)
// }

// func calculate(numbers ...int) float64 {
// 	var total int = 0
// 	for _, number := range numbers {
// 		total += number
// 	}

// 	var avg = float64(total) / float64(len(numbers))
// 	return avg
// }

/** 
	//  Penggunaan Fungsi fmt.Sprintf()
		Fungsi fmt.Sprintf() pada dasarnya sama dengan fmt.Printf(), hanya saja fungsi ini tidak menampilkan nilai, melainkan mengembalikan nilainya dalam bentuk string. Pada kasus di atas, nilai kembalian fmt.Sprintf() ditampung oleh variabel msg.

		Selain fmt.Sprintf(), ada juga fmt.Sprint() dan fmt.Sprintln().


	//	Penggunaan Fungsi float64()
		Sebelumnya sudah dibahas bahwa float64 merupakan tipe data. Tipe data jika ditulis sebagai fungsi (penandanya ada tanda kurungnya) berguna untuk casting. Casting sendiri adalah teknik untuk konversi tipe sebuah data ke tipe lain. Hampir semua jenis tipe data dasar yang telah dipelajari di bab 9 bisa digunakan untuk casting. Dan cara penerepannya juga sama, cukup panggil sebagai fungsi, lalu masukan data yang ingin dikonversi sebagai parameter.

		Pada contoh di atas, variabel total yang tipenya adalah int, dikonversi menjadi float64, begitu juga len(numbers) yang menghasilkan int dikonversi ke float64.

		Variabel avg perlu dijadikan float64 karena penghitungan rata-rata lebih sering menghasilkan nilai desimal.

		Operasi bilangan (perkalian, pembagian, dan lainnya) di Go hanya bisa dilakukan jika tipe datanya sejenis. Maka dari itulah perlu adanya casting ke tipe float64 pada tiap operand.

	//	Pengisian Parameter Fungsi Variadic Menggunakan Data Slice.
		var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
		var avg = calculate(numbers...)
		var msg = fmt.Sprintf("Rata-rata : %.2f", avg)

		fmt.Println(msg)
*/

//	Fungsi Dengan Parameter Biasa & Variadic
import "fmt"

import "strings"

func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}

// func main() {
//     yourHobbies("wick", "sleeping", "eating")
// }

func main() {
    var hobbies = []string{"sleeping", "eating"}
    yourHobbies("wick", hobbies...)
}