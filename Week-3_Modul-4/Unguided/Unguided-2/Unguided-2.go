package main 
import "fmt"
const Pi = 3.14

func hitungPersegi(sisi int) {
	var luas, keliling int
	luas = sisi * sisi
	keliling = 4 * sisi
	fmt.Println("Luas persegi: ", luas)
	fmt.Println("Keliling persegi: ", keliling)
}

func hitungPersegiPanjang(panjang, lebar int) {
	var luas, keliling int
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)
	fmt.Println("Luas persegi panjang: ", luas)
	fmt.Println("Keliling persegi panjang: ", keliling)
}

func hitungLingkaran(jarijari float64) {
	var luas, keliling float64
	luas = Pi * jarijari * jarijari
	keliling = 2 * Pi * jarijari
	fmt.Println("Luas lingkaran: ", luas)
	fmt.Println("Keliling lingkaran: ", keliling)
}

func main() {
	var pil int

	fmt.Println("--- PROGRAM BANGUN DATAR ---")
	fmt.Println("1. Hitung luas & keliling persegi")
	fmt.Println("2. Hitung luas & keliling persegi panjang")
	fmt.Println("3. Hitung luas & keliling lingkaran")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pil)
	fmt.Println()

	switch pil {
	case 1:
		var sisi int
		fmt.Print("Masukkan sisi: ")
		fmt.Scan(&sisi)
		hitungPersegi(sisi)
	case 2:
		var p, l int
		fmt.Print("Masukkan panjang: ")
		fmt.Scan(&p)
		fmt.Print("Masukkan lebar: ")
		fmt.Scan(&l)
		hitungPersegiPanjang(p, l)
	case 3:
		var r float64
		fmt.Print("Masukkan jari-jari: ")
		fmt.Scan(&r)
		hitungLingkaran(r)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}