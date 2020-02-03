package main

import "fmt"

func main()  {
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// Penggunaan Keyword for Dengan Argumen Hanya Kondisi
	var i = 0

	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}

	// Penggunaan Keyword for Tanpa Argumen
	var w = 0

	for {
		fmt.Println("Angka", w)

		w++
		if w == 5 {
			break
		}
	}

	// Penggunaan Keyword break & continue
	for q := 1; q <= 10; q++ {
		if q % 2 == 1 {
			continue
		}

		if q > 8 {
			break
		}

		fmt.Println("Angka", q)
	}

	// Perulangan Bersarang
	for r := 0; r < 5; r++ {
		for j := r; j < 5; j++ {
			fmt.Println(j, " ")
		}

		fmt.Println("")
	}

	// Pemanfaatan Label Dalam Perulangan
	outerLoop:

	for y := 0; y < 5; y++ {
		for u := 0; u < 5; u++ {
			if y == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", y, "][", u, "]", "\n")
		}
	}
}