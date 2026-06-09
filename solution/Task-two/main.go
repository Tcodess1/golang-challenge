package main

import "fmt"

func Q2IsEven(n int) bool {
	num  := n % 2
	if num == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(Q2IsEven(2))
}
