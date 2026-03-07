package main
import "fmt"
func main() {
	var tahun int
	var status bool

	fmt.Print("Tahun : ")
	fmt.Scan(&tahun)
	status = (tahun % 400 == 0) || (tahun % 4 == 0 && tahun % 100 != 0)
	fmt.Println("Kabisat :", status)
}
