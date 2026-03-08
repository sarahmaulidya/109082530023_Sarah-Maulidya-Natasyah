package main
import "fmt"
func main() {
	var angka, pilihan int

	fmt.Println("-- MENU --")
	fmt.Println("1. Cek angka == 10")
	fmt.Println("2. Cek genap ganjil")
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		fmt.Print("Masukkan angka: ")
		fmt.Scan(&angka)
		if angka == 10 {
			fmt.Println("Angka adalah 10")
		} else {
			fmt.Println("Angka bukan 10")
		}
	
	case 2:
	fmt.Print("Masukkan angka : ")
	fmt.Scan(&angka)
	if angka % 2 == 0 {
		fmt.Println("Angka genap")
	} else if angka % 2 != 0 {
		fmt.Println("Angka ganjil")
	}

	default:
		fmt.Println("Pilihan tidak tersedia")
	}
}