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

func cariKetua(suara arrSuara) int {
	ketua := 0
	for i := 1; i <= NMAX; i++ {
		if suara[i] > 0 && (ketua == 0 || suara[i] > suara[ketua]) {
			ketua = i
		}
	}
	return ketua
}

func cariWakil(suara arrSuara, ketua int) int {
	wakil := 0
	for i := 1; i <= NMAX; i++ {
		if suara[i] > 0 && i != ketua && (wakil == 0 || suara[i] > suara[wakil]) {
			wakil = i
		}
	}
	return wakil
}

func main() {
	var suara arrSuara
	var masuk, sah int
	var ketua, wakil int

	masuk, sah = bacaSuara(&suara)
	
	fmt.Println("Suara masuk:", masuk)
	fmt.Println("Suara sah:", sah)

	ketua = cariKetua(suara)
	wakil = cariWakil(suara, ketua)

	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)

}