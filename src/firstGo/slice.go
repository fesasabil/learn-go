package main

import "fmt"

func main() {
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0])

	/** 
		perbedaan antara array dan slice
		var fruitsA = []string{"apple", "grape"} // slice
		var fruitsB = [2]string{"apple", "melon"} // array
		var fruitsC = [...]string{"banana", "watermelon"} // array
	*/

	// Hubungan Slice Dengan Array & Operasi Slice
	var fruit = []string{"apple", "grape", "banana", "melon"}

	var newFruits = fruit[0:3]

	fmt.Println(newFruits)

	/** 
		fruits[0:2]	[apple, grape]	semua elemen mulai indeks ke-0, hingga sebelum indeks ke-2
		fruits[0:4]	[apple, grape, banana, melon]	semua elemen mulai indeks ke-0, hingga sebelum indeks ke-4
		fruits[0:0]	[]	menghasilkan slice kosong, karena tidak ada elemen sebelum indeks ke-0
		fruits[4:4]	[]	menghasilkan slice kosong, karena tidak ada elemen yang dimulai dari indeks ke-4
		fruits[4:0]	[]	error, pada penulisan fruits[a,b] nilai a harus lebih besar atau sama dengan b
		fruits[:]	[apple, grape, banana, melon]	semua elemen
		fruits[2:]	[banana, melon]	semua elemen mulai indeks ke-2
		fruits[:2]	[apple, grape]	semua elemen hingga sebelum indeks ke-2
	*/

	// Slice Merupakan Tipe Data Reference
	var fruitss = []string{"apple", "grape", "banana", "melon"}

	var aFruits = fruitss[0:3]
	var bFruits = fruitss[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruitss)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits)

	baFruits[0] = "pinnaple"

	fmt.Println(fruitss)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits)

	// Fungsi len()
	var fruitsD = []string{"potato", "strawbery", "blubery"}
	fmt.Println(len(fruitsD))

	//  Fungsi cap()
	var fruitA = []string{"apple", "grape", "banana", "melon"}

	fmt.Println(len(fruitA))
	fmt.Println(cap(fruitA))

	var abFruits = fruitA[0:3]
	fmt.Println(len(abFruits)) // len: 3
	fmt.Println(cap(abFruits)) // cap: 4

	var bbFruits = fruitA[1:4]
	fmt.Println(len(bbFruits)) // len: 3
	fmt.Println(cap(bbFruits)) // cap: 3

	/** 
		Kode			Output					len()	cap()
		fruits[0:4]		[buah buah buah buah]	4		4	
		aFruits[0:3]	[buah buah buah ----]	3		4
		bFruits[1:3]	---- [buah buah buah]	3		3
	*/

	// Fungsi append()
	var animal  = []string{"dog", "cat", "tiger"}
	var cAnimal = append(animal, "lion")

	fmt.Println(animal)
	fmt.Println(cAnimal)

	/** 
		Ada 2 hal yang perlu diketahui dalam penggunaan fungsi ini.

		- Ketika jumlah elemen dan lebar kapasitas adalah sama (len(fruits) == cap(fruits)), maka elemen baru hasil append() merupakan referensi baru.
		- Ketika jumlah elemen lebih kecil dibanding kapasitas (len(fruits) < cap(fruits)), elemen baru tersebut ditempatkan kedalam cakupan kapasitas, menjadikan semua elemen slice lain yang referensi-nya sama akan berubah nilainya.
	*/

	var fruitsss = []string{"apple", "grape", "banana"}
	var bFruitss = fruitsss[0:2]

	fmt.Println(cap(bFruitss)) // 3
	fmt.Println(len(bFruitss)) // 2

	fmt.Println(fruitsss)  // ["apple", "grape", "banana"]
	fmt.Println(bFruitss) // ["apple", "grape"]

	var cFruits = append(bFruitss, "papaya")

	fmt.Println(fruitsss)  // ["apple", "grape", "papaya"]
	fmt.Println(bFruitss) // ["apple", "grape"]
	fmt.Println(cFruits) // ["apple", "grape", "papaya"]
}