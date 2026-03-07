package main 
import "fmt"
func main() {
	var gram, kg, sisa, ongkir, tambahan, total int

	fmt.Print("Masukkan total berat (gram): ")
	fmt.Scan(&gram)

	kg = gram / 1000
	sisa = gram % 1000

	ongkir = 10000 * kg
	tambahan = 0

	if gram > 10000 {
		tambahan = 0
	} else if sisa >= 500 {
		tambahan = 5 * sisa
	} else if sisa <= 500 {
		tambahan = 15 * sisa
	} else {
		tambahan = 0
	}
	total = ongkir + tambahan

	fmt.Println("===== Detail Perhitungan =====")
	fmt.Println("Detail berat: ", kg, "kg + ", sisa, "gr")
	fmt.Println("Detail biaya: Rp", ongkir, "+ Rp", tambahan)
	fmt.Println("Total biaya: Rp", total)
}