package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MagicSearch(slice []int, value int) int {
	ini := 0
	fim := len(slice) -1

	for ini <= fim{

		meio := (ini+fim)/2

		if value == slice[meio]{
			for meio +1 < len(slice) && slice[meio+1] == value {
				meio++
			}
			return meio
		}

		if value > slice[meio]{
			ini = meio +1
		}
		if value < slice[meio]{
			fim = meio-1
		}
	}
	return ini
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	slice := make([]int, 0, 1)
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}

	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	result := MagicSearch(slice, value)
	fmt.Println(result)
}
