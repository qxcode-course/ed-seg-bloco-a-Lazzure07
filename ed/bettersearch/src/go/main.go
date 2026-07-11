package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BetterSearch(slice []int, value int) (bool, int) {
	inicio := 0
	fim := len(slice) - 1

	for inicio <= fim {
		meio := (inicio + fim) / 2

		if slice[meio] == value {
			return true, meio
		} else if inicio == fim && value < slice[meio] {
			return false, inicio
		} else if inicio == fim && value > slice[meio] {
			return false, inicio+1
		}

		if value < slice[meio] {
			fim = meio - 1
		} else {
			inicio = meio + 1
		}
	}

	return false, 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	slice := []int{}
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}
	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	found, result := BetterSearch(slice, value)
	if found {
		fmt.Println("V", result)
	} else {
		fmt.Println("F", result)
	}
}
