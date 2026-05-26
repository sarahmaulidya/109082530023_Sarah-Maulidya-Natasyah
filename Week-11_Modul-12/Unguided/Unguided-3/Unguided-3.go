package main 

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func main() {
	var n, k, hasil int

	fmt.Scan(&n, &k)

	isiArray(n)

	hasil = posisi(n, k)

	if hasil == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(hasil)
	}
}

func isiArray(n int) {
	i := 0
	for i < n {
		fmt.Scan(&data[i])
		i++
	}
}

func posisi(n, k int) int {
	found := -1
	kr := 0
	kn := n - 1

	for kr <= kn && found == -1 {
		med := (kr + kn) / 2

		if k < data[med] {
			kn = med - 1
		} else if k > data[med] {
			kr = med + 1
		} else {
			found = med
		}
	}
	return found
}