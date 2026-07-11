package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tostr(vet []int) string {
	if len(vet) == 0{
		return "[]"
	}

	s := "["

	for i, v := range vet{
		if i > 0{
			s += ", "
		}
		s += fmt.Sprintf("%v", v)
	}

	return s + "]"
}

func tostrrev(vet []int, inicio, fim int) string {
	var aux string
	if inicio >= fim{
		aux = tostr(vet)
		return aux
	}
	vet[inicio], vet[fim] = vet[fim], vet[inicio]
	inicio++
	fim--

	return tostrrev(vet, inicio, fim)
}

// reverse: inverte os elementos do slice
func reverse(vet []int, inicio, fim int) {
    if inicio >= fim{
        return
    }
    vet[inicio], vet[fim] = vet[fim], vet[inicio]
    inicio++
    fim--

    reverse(vet, inicio, fim)
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	soma := 0
	for _, v := range vet{
		soma += v
	}
	return soma
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	produto := 1
	for _, v := range vet{
		produto *= v
	}
	return produto
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {

	if len(vet) == 0{
		return -1
	}

	var rec func(v []int) (int, int)

	rec = func(v []int) (int, int) {

		if len(v) == 1 {
			return 0, v[0]
		}

		indice, valor := rec(v[1:])

		if v[0] < valor {
			return 0, v[0]
		}

		return indice + 1, valor
	}

	indice, _ := rec(vet)

	return indice
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet, 0, len(vet)-1))
		case "reverse":
			reverse(vet, 0, len(vet)-1)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
