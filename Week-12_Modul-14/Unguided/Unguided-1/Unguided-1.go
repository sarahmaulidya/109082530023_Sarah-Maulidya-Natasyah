package main

import "fmt"

const max = 1000000

type array [max]int

func selectionSort(arr *array, n int) {
	var idx_min, temp int

	for i := 0; i < n - 1; i++ {
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

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		var data array
		for j := 0; j < m; j++ {
			var no int
			fmt.Scan(&no)
			data[j] = no
		}

		selectionSort(&data, m)

		for j := 0; j < m; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(data[j])
		}
		fmt.Println()
	}
}