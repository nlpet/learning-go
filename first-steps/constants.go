package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {

	fmt.Println(s)

	const n = 500
	const d = 3e19 / n

	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
