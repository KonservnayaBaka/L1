package main

import (
	"fmt"
	"math/big"
)

func main() {
	aStr := "92233720368547758071" //+01 от размера int
	bStr := "92233720368547758072" //+2 от размера int

	a := new(big.Int)
	if _, ok := a.SetString(aStr, 10); !ok {
		panic("invalid a: " + aStr)
	}
	b := new(big.Int)
	if _, ok := b.SetString(bStr, 10); !ok {
		panic("invalid b: " + bStr)
	}

	sum := new(big.Int).Add(a, b)
	sub := new(big.Int).Sub(a, b)
	mul := new(big.Int).Add(a, b)
	div := new(big.Int).Div(a, b)

	fmt.Printf("Sum: %s\nSubtraction: %s\nMultiplication: %s\nDivision: %s\n", sum.String(), sub.String(), mul.String(), div.String())
}
