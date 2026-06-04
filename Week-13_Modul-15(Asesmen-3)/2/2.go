package main

import "fmt"

type pemain struct {
	nama string
	gol int
	assist int
}

const NMAX = 1000

type arrPemain [NMAX]pemain

func main() {
	var daftar arrPemain
	var  n int
	var ndepan, nbelakang string

	fmt.Println("Masukkan Data input: ")
	fmt.Scan(&n)
	for i := 0; i < n && i < NMAX; i++ {
		fmt.Scan(&ndepan, &nbelakang, &daftar[i].gol, &daftar[i].assist)
		daftar[i].nama = ndepan + " " + nbelakang
	}
	
	var posisi, idx_max int
	var temp pemain

	for posisi = 1; posisi <= n - 1; posisi++ {
		idx_max = posisi - 1
		for i := posisi; i < n; i++ {
			if daftar[idx_max].gol < daftar[i].gol || (daftar[idx_max].gol == daftar[i].gol && daftar[idx_max].assist <= daftar[i].assist) {
				idx_max = i
			}
		}
		temp = daftar[idx_max]
		daftar[idx_max] = daftar[posisi - 1]
		daftar[posisi - 1] = temp
	}

	fmt.Println("\nHasil Sorting: ")
	for i := 0; i < n; i++ {
		fmt.Println(daftar[i].nama, daftar[i].gol, daftar[i].assist)
	}
}