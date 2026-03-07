package main 
import "fmt"
func main() {
	var angka int
	
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	for angka < 10 {
		angka++
		fmt.Println("Angka sekarang ", angka)
	}
}