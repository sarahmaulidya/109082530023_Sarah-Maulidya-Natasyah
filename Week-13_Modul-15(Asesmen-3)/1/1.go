package main

import "fmt"

const NMAX = 1000000

type arrInt [NMAX]int

func sorting(T *arrInt, n int) {
	var idx_min, temp int

	for i := 1; i <= n - 1; i++ {
		idx_min = i - 1
		for j := i; j <= n - 1; j++ {
			if T[idx_min] > T[j] {
				idx_min = j
			}
		}
		temp = T[idx_min]
		T[idx_min] = T[i - 1]
		T[i - 1] = temp
	}
}

func median(T arrInt, n int) float64 {
	var mid int 

	mid = n / 2

	if n % 2 == 0 {
		return float64(T[mid - 1] + T[mid]) / 2.0
	} else {
		return float64(T[mid])
	}
}

func main() {
	var data arrInt
	var bil, n int

	fmt.Println("Input data masukan: ")
	fmt.Scan(&bil)

	n = 0

	for bil != -5313541 && n < NMAX {
		if bil == 0 {
			sorting(&data, n)
			fmt.Println("Median: ")
			fmt.Println(median(data, n))
		} else {
			data[n] = bil
			n++
		}
		fmt.Scan(&bil)
	}
}