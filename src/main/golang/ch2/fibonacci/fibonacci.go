package fibonacci

import "math/big"

func Fibonacci(n uint) *big.Int {
	a := big.NewInt(0)
	b := big.NewInt(1)
	if n == 0 {
		return a
	}
	if n == 1 {
		return b
	}
	for range n - 1 {
		a, b = b, a.Add(a, b)
	}
	return b
}
