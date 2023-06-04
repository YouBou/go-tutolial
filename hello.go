package main

import (
	"fmt"
	"strconv"

	"github.com/YouBou/go-tutolial/hello"
)

type intp int

func (num intp) IsPrime() bool {
	n := int(num)
	for i := 2; i <= (n / 2); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func (num intp) PrimeFactor() []int {
	ar := []int{}
	x := int(num)
	n := 2
	for x > n {
		if x%n == 0 {
			x /= n
			ar = append(ar, n)
			continue
		}
		if n == 2 {
			n++
			continue
		}
		n += 2
	}
	ar = append(ar, x)
	return ar
}

func (num *intp) doPrime() {
	// 引数numにはintpのポインタが渡される

	// PrimeFactor()で素因数分解されたスライスを取得
	pf := num.PrimeFactor()
	// 取得したスライスの最後の値をnumの値に設定
	*num = intp(pf[len(pf)-1])
}

func main() {
	s := hello.Input("type a number")
	n, _ := strconv.Atoi(s)
	x := intp(n)
	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrimeFactor())
	x.doPrime()
	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrimeFactor())
	x++
	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrimeFactor())
}
