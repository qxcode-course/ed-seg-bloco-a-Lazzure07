package main

import "fmt"

// x: número que está sendo testado
// div: divisor que está sendo testado
func eh_primo(x, div int) bool {

	if x == 2{
		return true
	}

	if  x < 2 || x%2 == 0{
		return false
	}

	if div*div>x{
		return true
	}

	if x%div != 0{
		div += 2
		return eh_primo(x, div)
	}

	return false;
}

func main() {
	var x int
	fmt.Scan(&x)
	if eh_primo(x, 3) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
