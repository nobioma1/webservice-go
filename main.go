package main

import (
	"fmt"
)

func main() {
	// variable declarations

	// Verbose declaration
	var i int // decalaration and type
	i = 10    // setting to an integer
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	// implicit initialization syntax
	isTrue := false
	fmt.Println(isTrue)

	// complex data type
	result := complex(2, 3)
	fmt.Println(result)

	// single line multiple assignment
	r, im := real(result), imag(result) // pull out the real and imaginary numbers
	fmt.Println(r, im)
}
