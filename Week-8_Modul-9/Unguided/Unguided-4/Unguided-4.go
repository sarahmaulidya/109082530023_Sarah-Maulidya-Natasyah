package main 

import "fmt"

const nmax int = 127

type tabel [nmax]rune

func isiArray(data *tabel, count *int) {
	*count = 0
	for {
		var input string
		fmt.Scan(&input)

		if input == "." || *count >= nmax {
			break
		}

		(*data)[*count] = rune(input[0])
		(*count)++
	}
}	

func cetakArray(data tabel, count int) {
	for i := 0; i < count; i++ {
		fmt.Print(string(data[i]), " ")
	}
	fmt.Println()
}

func balikArray(data *tabel, count int) {
	for i := 0; i < count / 2; i++ {
		tempValue := (*data) [i]
		(*data)[i] = (*data) [count - 1 - i]
		(*data) [count - 1 - i] = tempValue
	}
}

func Palindrom(data tabel, count int) bool {
	dataOri := data

	balikArray(&data, count)

	for i := 0; i < count; i++ {
		if dataOri[i] != data[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int

	fmt.Print("Input teks: ")
	isiArray(&tab, &n)

	var tabBalik tabel
	for i := 0; i < n; i++ {
		tabBalik[i] = tab[i]
	}

	balikArray(&tabBalik, n)

	fmt.Print("Teks terbalik: ")
	cetakArray(tabBalik, n)

	iniPalindrom := Palindrom(tab, n)

	fmt.Print("Palindrom: ")
	if iniPalindrom {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}