package main

import (
	"fmt"
	"math/big"
)

func main() {
	/*Для работы с большими числами воспользуемся пакетом big*/
	a, b := big.NewInt(2<<21), big.NewInt(2<<20)

	var sum big.Int
	sum.Add(a, b)

	var sub big.Int
	sub.Sub(a, b)

	var mul big.Int
	mul.Mul(a, b)

	var div big.Int
	div.Div(a, b)

	fmt.Println("+: ", sum.Int64(), " -: ", sub.Int64(), " *: ", mul.Int64(), "/: ", div.Int64())
}
