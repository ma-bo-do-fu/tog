package main

import "fmt"

func fibonacci() func() int {
	fibo, tmp := 0, 1
	return func() int {
		tmp, fibo = fibo, tmp+fibo
		return tmp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

/*
0
1
1
2
3
5
8
13
21
34
 */
