package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}
	fmt.Printf("%v\n", p) // instance of point struct

	fmt.Printf("%+v\n", p) // include the struct's field names
	fmt.Printf("%#v\n", p) // syntax representation of the value

	fmt.Printf("%T\n", p)    // type of a value
	fmt.Printf("%t\n", true) // format a bool

	fmt.Printf("%d\n", 5331) // format an integer, base 10
	fmt.Printf("%b\n", 42)   // binary representation
	fmt.Printf("%c\n", 42)   // character corresponding to integer
	fmt.Printf("%x\n", 345)  // hex encoding

	fmt.Printf("%e\n", 3561.0345)
	fmt.Printf("%E\n", 3561.0345)

	fmt.Printf("%s\n", "\"string\"")
	fmt.Printf("%q\n", "\"string\"")
	fmt.Printf("%x\n", "hex this")

	fmt.Printf("%p\n", &p) // pointer

	// specify width and justification
	fmt.Printf("|%5d|%5d|\n", 12, 345)
	fmt.Printf("|%5.2f|%5.2f|\n", 1.2, 3.45)
	fmt.Printf("|%-5.2f|%-5.2f|\n", 1.2, 3.45)
	fmt.Printf("|%5s|%5s|\n", "foo", "b")

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
