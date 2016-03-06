package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func sumUpTo(n int) int {
	if n == 0 {
		return 0
	}
	return n + sumUpTo(n-1)
}

func main() {
	fmt.Println(fact(5))
	fmt.Println(sumUpTo(5))
}
