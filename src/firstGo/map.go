package main

import "fmt"

func main()  {
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"])
	fmt.Println("mei", chicken["mei"])

	// // Inisialisasi Nilai Map
	// var data map[string]int
	// data["one"] = 1

	// data = map[string]int{}
	// data["one"] = 1

	// // cara vertikal
	// var chicken1 = map[string]int{"januari": 50, "februari": 40}

	// // cara horizontal
	// var chicken2 = map[string]int{
	// 	"januari": 50,
	// 	"februari": 40,
	// }

	// var chicken3 = map[string]int{}
	// var chicken4 = make(map[string]int)
	// var chicken5 = *new(map[string]int)

	// Iterasi Item Map Menggunakan for - range
	var chickens = map[string]int{
		"januari": 50,
		"februari": 40,
		"maret": 34,
		"april": 67,
	}

	for key, val := range chickens {
		fmt.Println(key, " \t:", val)
	}

	// Menghapus Item Map
	var chicken3 = map[string]int{"januari": 50, "februari": 40}

	fmt.Println(len(chicken3))
	fmt.Println(chicken3)

	delete(chicken3, "januari")

	fmt.Println(len(chicken3))
	fmt.Println(chicken3)

	// Deteksi Keberadaan Item Dengan Key Tertentu
	var chicken4 = map[string]int{"januari": 50, "februari": 40}
	var value, isExist = chicken4["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	//  Kombinasi Slice & Map
	var foods = []map[string]string{
		map[string]string{"name": "hotdog", "price": "murah"},
		map[string]string{"name": "burger", "price": "murah"},
		map[string]string{"name": "KFC", "price": "mahal"},
	}

	for _, food := range foods {
		fmt.Println(food["price"], food["name"])
	}
}