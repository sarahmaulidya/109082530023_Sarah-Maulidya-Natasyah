package main

import 	"fmt"

type titik struct {
	x, y int
}

type lingkaran struct {
	pusat titik
	radius int
}

func cekPosisi(t titik, l lingkaran) bool {
	jarakX := t.x - l.pusat.x
	jarakY := t.y - l.pusat.y
	return (jarakX * jarakX + jarakY * jarakY) <= (l.radius * l.radius)
}

func main() {
	var l1, l2 lingkaran
	var titikUji titik

	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.radius)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.radius)
	fmt.Scan(&titikUji.x, &titikUji.y)

	inL1 := cekPosisi(titikUji, l1)
	inL2 := cekPosisi(titikUji, l2)

	if inL1 && inL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}