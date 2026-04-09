package main
import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int{
	if jumlahHari > 3 && tujuan == "domestik"{
		jumlahHari = 3
	} else if jumlahHari == 8 && tujuan == "mancanegara"{
		jumlahHari = 8
	}
	return jumlahHari
}

func biayaPerHari(jumlahMhs int) int{
	return jumlahMhs * (70000 + 250000 + 300000)
}

func perhitunganBiaya(jumlahMhs, lamaPerjalanan int, tujuan string, totalBiaya float64) {
	
}

func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64

	fmt.Print("Jumlah mahasiswa: ")
	fmt.Scan(&jumlah)

	fmt.Print("Masukkan lama hari study tour: ")
	fmt.Scan(&lama)

	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara): ")
	fmt.Scan(&tujuan)
	fmt.Println()
	
	fmt.Print("Biaya perjalanan yang harus dikeluarkan Tel-U: ",biaya)
}