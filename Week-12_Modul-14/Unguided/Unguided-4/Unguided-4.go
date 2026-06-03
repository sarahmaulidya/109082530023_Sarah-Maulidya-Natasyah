package main

import "fmt"

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit, &pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

func CetakTerFavorit(pustaka DaftarBuku, n int) {
	idx_max := 0
	for i := 1; i < n; i++ {
		if pustaka[i].rating > pustaka[idx_max].rating {
			idx_max = i
		}
	}
	fmt.Println(pustaka[idx_max].judul, pustaka[idx_max].penulis,
		pustaka[idx_max].penerbit, pustaka[idx_max].tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var temp Buku

	for i := 1; i <= n-1; i++ {
		j := i
		temp = pustaka[j]
		for j > 0 && temp.rating > pustaka[j-1].rating {
			pustaka[j] = pustaka[j-1]
			j--
		}
		pustaka[j] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	var batas int

	batas = 5
	if n < 5 {
		batas = n
	}

	for i := 0; i < batas; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(pustaka[i].judul)
	}
	fmt.Println()
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	var left, right, found int

	left = 0
	right = n - 1
	found = -1

	for left <= right {
		mid := (left + right) / 2
		if pustaka[mid].rating == r {
			found = mid
			break
		} else if pustaka[mid].rating > r {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if found != -1 {
		fmt.Println(pustaka[found].judul, pustaka[found].penulis, pustaka[found].penerbit, pustaka[found].tahun, pustaka[found].rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var n int
	var pustaka DaftarBuku

	fmt.Scan(&n)
	DaftarkanBuku(&pustaka, n)
	CetakTerFavorit(pustaka, n)
	UrutBuku(&pustaka, n)
	Cetak5Terbaru(pustaka, n)

	var cariRating int
	fmt.Scan(&cariRating)
	CariBuku(pustaka, n, cariRating)
}
