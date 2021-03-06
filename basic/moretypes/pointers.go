/*
Pointers
Goはポインタを扱います。 ポインタは値のメモリアドレスを指します。

変数 T のポインタは、 *T 型で、ゼロ値は nil です。

var p *int
& オペレータは、そのオペランド( operand )へのポインタを引き出します。

i := 42
p = &i
* オペレータは、ポインタの指す先の変数を示します。

fmt.Println(*p) // ポインタpを通してiから値を読みだす
*p = 21         // ポインタpを通してiへ値を代入する
これは "dereferencing" または "indirecting" としてよく知られています。

なお、C言語とは異なり、ポインタ演算はありません。
*/

package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

}

/*
42
21
73
*/
