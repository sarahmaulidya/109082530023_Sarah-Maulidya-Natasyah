package main

import "fmt"

const NMAX = 1000000

type partai struct {
	nama int
	suara int
}

type tabPartai [NMAX]partai

func main(){
	var p tabPartai
	var n, kode, idx int

	n = 0

	fmt.Println("Masukkan proses input suara: ")
	fmt.Scan(&kode)
	
	for kode != -1 {
		idx = posisi(p, n, kode)
		if idx == -1 {
			p[n].nama = kode
			p[n].suara = 1
			n++
		}else{
			p[idx].suara++
		}
		fmt.Scan(&kode)
	}

	var temp partai
	
	for i := 1; i <= n-1; i++ {
		j := i
		temp = p[j]
		for j > 0 && temp.suara > p[j-1].suara {
			p[j] = p[j-1]
			j--
		}
		p[j] = temp
	}

	fmt.Println("\nHasil Perhitungan suara: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%v(%v) ", p[i].nama, p[i].suara)
	
	}
	fmt.Println()
}

func posisi(t tabPartai, n int, nama int) int {
	var i, ketemu int
	
	i = 0
	ketemu = -1
	for i < n && ketemu == -1 {
		if t[i].nama == nama {
			ketemu = i
		}
		i++
	}
	return ketemu
}