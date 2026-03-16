package main 
import "fmt"

func hitungBiaya(jenis string, jamMasuk, jamKeluar int) int {
	var biayaTotal int
	biayaTotal = 0

	for j := jamMasuk; j < jamKeluar; j++ {
		if jenis == "motor" {
			if j < 17 || j == 24 {
				biayaTotal += 4000
			} else {
				biayaTotal += 5000
			}
		} else if jenis == "mobil" {
			if j < 17 || j == 24 {
				biayaTotal += 6000
			} else {
				biayaTotal += 7000
			}
		}
	}
	return biayaTotal
}

func main() {
	var jenis string
	var masuk, keluar, biaya, noKendaraan, totalPendapatan int

	noKendaraan = 1

	fmt.Println("=== Rekap Tarif Parkir Cafe Per Hari ===")

	for {
		fmt.Printf("\n*Kendaraan %d\n", noKendaraan)
		fmt.Print("Kendaraan apa? (motor/mobil/- untuk selesai): ")
		fmt.Scan(&jenis)

		if jenis == "-" {
			break
		}

		fmt.Print("Masukkan Jam Masuk Kendaraan (0-24): ")
		fmt.Scan(&masuk)
		fmt.Print("Masukkan Jam Keluar Kendaraan (0-24): ")
		fmt.Scan(&keluar)

		biaya = hitungBiaya(jenis, masuk, keluar)
		fmt.Printf("Biaya parkir %s %d: %d\n\n", jenis, noKendaraan, biaya)
		fmt.Println("=========================")
		totalPendapatan += biaya

		noKendaraan++
	}

	fmt.Printf("\n*** Total Pendapatan Hari Ini Adalah %d ***\n", totalPendapatan)
}