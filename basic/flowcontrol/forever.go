package main

import (
	"fmt"
)

func main() {
	i := 1
	for {
		if (i > 1000) {
			break
		}
		fmt.Println(i)
		i++
	}
}
