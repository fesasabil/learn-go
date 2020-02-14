package main

import "fmt"
import "strings"


/** 
	Cara membuat parameter fungsi adalah dengan langsung menuliskan skema fungsi nya sebagai tipe data. Contohnya bisa dilihat pada kode berikut.

	// Alias Skema Closure
		closure bisa dimanfaatkan sebagai tipe parameter, contohnya seperti pada fungsi filter(). Pada fungsi tersebut kebetulan skema tipe parameter closure-nya tidak terlalu panjang, hanya ada satu buah parameter dan satu buah nilai balik.

		pada fungsi yang skema-nya cukup panjang, akan lebih baik jika menggunakan alias, apalagi ketika ada parameter fungsi lain yang juga menggunakan skema yang sama.Membuat alias fungsi berarti menjadikan skema fungsi tersebut menjadi tipe data baru. caranya dengan menggunakan keyword TYPE. contoh :

			type FilterCallback func(string) bool

			func filter(data []string, callback FilterCallback) []string {
				// ...
			}

		Skema func(string) bool diubah menjadi tipe dengan nama FilterCallback. Tipe tersebut kemudian digunakan sebagai tipe data parameter callback.

	// Penggunaan Fungsi string.Contains()
		Inti dari fungsi ini adalah untuk deteksi apakah sebuah substring adalah bagian dari string, jika iya maka akan bernilai TRUE, dan sebaliknya. Contoh penggunaannya :
			
			var result = string.Contains("Golang", "ang")
			// true
*/

func filter(data []string, callback func(string) bool) []string  {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

func main()  {
	var data = []string{"wick", "jason", "ethan"}
	var dataContainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	var dataLenght5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data asli \t\t:", data) // data asli : [wick jason ethan]

	fmt.Println("filter ada huruf \"o\"\t:", dataContainsO) // filter ada huruf "o" : [jason]

	fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5) // filter jumlah huruf "5" : [jason, ethan]
}
