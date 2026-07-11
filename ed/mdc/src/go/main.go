package main

import (
	"fmt"
	"math"
)

func mdc(a, b int) int {
	if a == 0 || b == 0{
		return int(math.Abs(  float64(a-b)  ))
	}
	R := a%b
	return mdc(b, R)
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(mdc(a, b))
}
