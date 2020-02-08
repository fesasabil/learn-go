// package main

// import "fmt"
// import "strings"

// func main() {
// 	var names = []string{"John", "Wick"}
// 	printMessage("Hallo", names)
// }

// func printMessage(message string, arr []string) {
// 	var nameString = strings.Join(arr, " ")
// 	fmt.Println(message, nameString)
// }

// Fungsi Dengan Return Value / Nilai Balik
// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func main() {
// 	rand.Seed(time.Now().Unix())
// 	var randomValue int

// 	randomValue = randomWithRange(2, 10)
// 	fmt.Println("random number:", randomValue)
// 	randomValue = randomWithRange(2, 10)
// 	fmt.Println("random number:", randomValue)
// 	randomValue = randomWithRange(2, 10)
// 	fmt.Println("random number:", randomValue)
// }

// func randomWithRange(min, max int) int {
// 	var value = rand.Int() % (max - min + 1) + min
// 	return value
// }

// Penggunaan Fungsi rand.Seed()
// rand.Seed(time.Now().Unix()) 
/** 
	- Fungsi ini diperlukan untuk memastikan bahwa angka random yang akan di-generate benar-benar acak. Kita bisa gunakan angka apa saja sebagai nilai parameter fungsi ini.

	- Fungsi rand.Seed() berada dalam package math/rand, yang harus di-import terlebih dahulu sebelum bisa dimanfaatkan.

	- Package time juga perlu di-import karena kita menggunakan fungsi (time.Now().Unix()) disitu.
*/

/** 
	//  Import Banyak Package
	import "fmt"
	import "math/rand"
	import "time"

	// atau

	import (
		"fmt"
		"math/rand"
		"time"
	)
*/

/** 
	// Deklarasi Parameter Bertipe Data Sama
	func nameOfFunc(paramA type, paramB type, paramC type) returnType
	func nameOfFunc(paramA, paramB, paramC type) returnType

	func randomWithRange(min int, max int) int
	func randomWithRange(min, max int) int
*/

// Penggunaan Keyword return Untuk Menghentikan Proses Dalam Fungsi
package main

import "fmt"

func main() {
    divideNumber(10, 2)
    divideNumber(4, 0)
    divideNumber(8, -4)
}

func divideNumber(m, n int) {
    if n == 0 {
        fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
        return
    }

    var res = m / n
    fmt.Printf("%d / %d = %d\n", m, n, res)
}
