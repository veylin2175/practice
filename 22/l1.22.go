package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(123456789)
	b := big.NewInt(123456789)

	fmt.Println(sum(a, b))
	fmt.Println(sub(a, b))
	fmt.Println(mul(a, b))
	fmt.Println(div(a, b))
}

func sum(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func sub(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func mul(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func div(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}
