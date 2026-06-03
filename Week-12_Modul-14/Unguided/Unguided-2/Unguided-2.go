package main

import "fmt"

func selectionSortA(arr []int, n int) {
	var idx_min, temp int

	for i := 0; i < n-1; i++ {
		idx_min = i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[idx_min] {
				idx_min = j
			}
		}
		temp = arr[idx_min]
		arr[idx_min] = arr[i]
		arr[i] = temp
	}
}

func SelectionSortD(arr []int, n int) {
	var idx_max, temp int

	for i := 0; i < n-1; i++ {
		idx_max = i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[idx_max] {
				idx_max = j
			}
		}
		temp = arr[idx_max]
		arr[idx_max] = arr[i]
		arr[i] = temp
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		var ganjil []int
		var genap []int

		for j := 0; j < m; j++ {
			var no int
			fmt.Scan(&no)

			if no%2 != 0 {
				ganjil = append(ganjil, no)
			} else {
				genap = append(genap, no)
			}
		}

		selectionSortA(ganjil, len(ganjil))
		SelectionSortD(genap, len(genap))

		first := true

		for j := 0; j < len(ganjil); j++ {
			if !first {
				fmt.Print(" ")
			}
			fmt.Print(ganjil[j])
			first = false
		}

		for j := 0; j < len(genap); j++ {
			if !first {
				fmt.Print(" ")
			}

			fmt.Print(genap[j])
			first = false
		}
		fmt.Println()
	}
}
