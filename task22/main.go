package main

import (
	"fmt"
	"math/big"
)

func sumBigInt(a, b string) (*big.Int, error) {
	x, y, ok := new(big.Int), new(big.Int), false

	x, ok = x.SetString(a, 10)
	if !ok {
		return nil, fmt.Errorf("error a str conv to bigInt")
	}
	y, ok = y.SetString(b, 10)
	if !ok {
		return nil, fmt.Errorf("error b str conv to bigInt")
	}
	return x.Add(x, y), nil
}

func subBigInt(a, b string) (*big.Int, error) {
	x, y, ok := new(big.Int), new(big.Int), false

	x, ok = x.SetString(a, 10)
	if !ok {
		return nil, fmt.Errorf("error a str conv to bigInt")
	}
	y, ok = y.SetString(b, 10)
	if !ok {
		return nil, fmt.Errorf("error b str conv to bigInt")
	}
	return x.Sub(x, y), nil
}

func multBigInt(a, b string) (*big.Int, error) {
	x, y, ok := new(big.Int), new(big.Int), false

	x, ok = x.SetString(a, 10)
	if !ok {
		return nil, fmt.Errorf("error a str conv to bigInt")
	}
	y, ok = y.SetString(b, 10)
	if !ok {
		return nil, fmt.Errorf("error b str conv to bigInt")
	}
	return x.Mul(x, y), nil
}

func divBigInt(a, b string) (*big.Int, error) {
	x, y, ok := new(big.Int), new(big.Int), false

	x, ok = x.SetString(a, 10)
	if !ok {
		return nil, fmt.Errorf("error a str conv to bigInt")
	}
	y, ok = y.SetString(b, 10)
	if !ok {
		return nil, fmt.Errorf("error b str conv to bigInt")
	}

	if y.Cmp(big.NewInt(0)) == 0 {
		return nil, fmt.Errorf("error division by zero")
	}

	return x.Div(x, y), nil
}

func main() {
	a, b := "7325290526723952356235", "32757295235862352"

	v, err := sumBigInt(a, b)
	fmt.Println("+", v, err)

	v, err = subBigInt(a, b)
	fmt.Println("-", v, err)

	v, err = multBigInt(a, b)
	fmt.Println("*", v, err)

	v, err = divBigInt(a, b)
	fmt.Println("/", v, err)

	v, err = divBigInt(a, "0")
	fmt.Println("/", v, err)
}
