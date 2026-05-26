package main

import "fmt"

const NMAX = 20

type arrSuara [NMAX + 1]int

func valid(x int) bool{
	return x >= 1 && x <= NMAX
}

func bacaSuara(suara *arrSuara) (masuk, sah int) {
	for {
		var x int

		fmt.Scan(&x)
		if x == 0 {
			break
		}
		masuk++

		if valid(x) {
			sah++
			suara[x]++
		}
	}
	return
}

func cetakHasil(suara arrSuara) {
	for i := 1; i <= NMAX; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}

	}
}

func main() {
	var suara arrSuara
	var masuk, sah int

	masuk, sah = bacaSuara(&suara)
	
	fmt.Println("Suara masuk:", masuk)
	fmt.Println("Suara sah:", sah)

	cetakHasil(suara)
}