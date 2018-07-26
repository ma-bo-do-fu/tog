//structのフィールドは、ドット( . )を用いてアクセスします。

package main

import "fmt"

type Vertex struct {
	Y int
	X int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

//4