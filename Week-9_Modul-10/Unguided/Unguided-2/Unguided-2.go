package main 

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	var berat [1000]float64

	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	var hasil [1000] float64
	jumlahWadah := 0

	i := 0
	for i < x {
		total := 0.0

		for j := 0; j < y && i < x; j++ {
			total += berat[i]
			i++
		}

		hasil[jumlahWadah] = total
		jumlahWadah++
	}

	sum := 0.0
	for i := 0; i < jumlahWadah; i++ {
		sum += hasil[i]
	}

	for i := 0; i < jumlahWadah; i++ {
		fmt.Print(hasil[i], " ")
	}
	fmt.Println()

	rata := sum / float64(jumlahWadah)

	fmt.Println(rata)
}