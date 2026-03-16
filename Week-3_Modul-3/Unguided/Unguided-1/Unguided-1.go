package main
import "fmt"

func factorial(n int) int {
	var hasil int
	if n <= 1 {
		return 1
	}

	hasil = 1
	for i := 2; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int
	var p1, p2, c1, c2 int
	fmt.Scan(&a, &b, &c, &d)

	p1 = permutation(a, c)
	c1 = combination(a, c)
	fmt.Printf("%d %d\n", p1, c1)

	p2 = permutation(b, d)
	c2 = combination(b, d)
	fmt.Printf("%d %d\n", p2, c2)
}