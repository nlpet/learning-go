package main

import (
	"fmt"
	"sort"
)

func main() {

	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 4, 3, 1}
	sort.Ints(ints)
	fmt.Println("Ints: ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

	floats := []float64{1.13, 1.14, 1.09, 1.04}
	sort.Float64s(floats)
	fmt.Println("Floats: ", floats)
}
