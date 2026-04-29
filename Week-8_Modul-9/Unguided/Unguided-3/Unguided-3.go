package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var pemenang [1000] string
	var totalLaga int

	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	lagaKe := 1
	for {
		fmt.Printf("Pertandingan %d: ", lagaKe)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang[totalLaga] = klubA
		} else if skorB > skorA {
			pemenang[totalLaga] = klubB
		} else {
			pemenang[totalLaga] = "Draw"
		}
		totalLaga++
	}

	fmt.Println()
	
	for i := 0; i < totalLaga; i++ {
		fmt.Printf("Hasil %d: %s\n", i+1, pemenang[i])
	}
	fmt.Println("Pertandingan selesai")
}