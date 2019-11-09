package main

import "fmt"

func main() {
	// Constant can be typed or untyped

	// Untyped Constants
	const ui = 12345
	const uf = 3.142592

	fmt.Println(ui)
	fmt.Println(uf)

	const ti int = 12345
	const tf float64 = 3.141592

	fmt.Println(ti)
	fmt.Println(tf)

	var answer = 3 * 0.333
	fmt.Println(answer)

	// Constant third will be of kind floating point
	const third = 1 / 3.0

	fmt.Println(third)

	// Constant zero will be of kind integer
	const zero = 1 / 3

	fmt.Println(zero)

	// must have like types to perform math
	const one int8 = 1
	const two = 2 * one // int8(2) * int8(1)

	fmt.Println(one)
	fmt.Println(two)

	// Max integer value on 64 bit architecture
	const maxInt = 9223372036854775807

	fmt.Println(maxInt)

	const (
		A1 = iota
		B1 = iota
		C1 = iota
	)

	fmt.Println("1:", A1, B1, C1)

	const (
		A2 = iota
		B2
		C2
	)

	fmt.Println("2:", A2, B2, C2)

	const (
		A3 = iota + 1
		B3
		C3
	)

	fmt.Println("3:", A3, B3, C3)

	const (
		Ldate		= 1 << iota
		Ltime
		Lmicroseconds
		Llongfile
		Lshortfile
		LUTC
	)

	fmt.Println("Log:", Ldate, Ltime, Lmicroseconds, Llongfile, Lshortfile, LUTC)
}
