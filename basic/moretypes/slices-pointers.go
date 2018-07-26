package main

import "fmt"

func main() {
	names := [4]string{
		"Jhon",
		"Paul",
		"Geroge",
		"Ringo",
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

/**
[Jhon Paul Geroge Ringo]
[Jhon Paul] [Paul Geroge]
[Jhon XXX] [XXX Geroge]
[Jhon XXX Geroge Ringo]
 */
