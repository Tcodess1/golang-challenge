package main

import "fmt"

func Q3Factorial(n int) int {
	num := 1

	for i := 2; i <= n; i++ {
		num *= i
	}
	return num
}

func main() {
	fmt.Println(Q3Factorial(4))
}