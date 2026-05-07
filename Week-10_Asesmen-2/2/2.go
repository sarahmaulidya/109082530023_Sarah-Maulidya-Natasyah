package main 

import "fmt"

const nMax = 51

type mahasiswa struct {
	NIM string
	nama string
	nilai int
}

type arrayMahasiswa [nMax] mahasiswa

func nilaiPertama(data arrayMahasiswa, N int, nim string) int {
	for i := 0; i < N; i++ {
		if data[i].NIM == nim {
			return data[i].nilai
		}
	}
	return -1
}

func nilaiTerbesar(data arrayMahasiswa, N int, nim string) int {
	nilaiMax := -1
	for i := 0; i < N; i++ {
		if data[i].NIM == nim {
			if data[i].nilai > nilaiMax {
				nilaiMax = data[i].nilai
			}
		}
	}
	return nilaiMax
}

func main() {
	var data arrayMahasiswa
	var N int
	fmt.Print("Masukkan jumlah data: ")
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan jumlah data ke-%d: ", i + 1)
		fmt.Scan(&data[i].NIM, &data[i].nama, &data[i].nilai)
	}

	var cariNIM string
	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya: ")
	fmt.Scan(&cariNIM)

	nilaiPertama := nilaiPertama(data, N, cariNIM)
	nilaiTerbesar := nilaiTerbesar(data, N, cariNIM)

	fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", cariNIM, nilaiPertama)
	fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", cariNIM, nilaiTerbesar)
}
