package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	var vetMen []int
	for _, v := range vet{
		if v > 0{
			vetMen = append(vetMen, v)
		}
	}
	return vetMen
}

func getCalmWomen(vet []int) []int {
	var vetWomen []int
	for _, v := range vet{
		if v < 0 && v > -10{
			vetWomen = append(vetWomen, v)
		}
	}
	return vetWomen
}

func sortVet(vet []int) []int {
	var vetSort []int
	
	for len(vet) > 0 {
		menor := vet[0]
		indice := 0
		for i, v := range vet{
			if v <= menor{
				menor = v
				indice = i
			}
		}
		vetSort = append(vetSort, menor)
		vet = append(vet[:indice], vet[indice+1:]...)
	}
	return vetSort
}

func sortStress(vet []int) []int {
	var vetAbs, vetSort []int
	var indice int

	for _, v := range vet{
		if v < 0{
			vetAbs = append(vetAbs, -v)
		}else{
			vetAbs = append(vetAbs, v)
		}
	}
	for len(vet) > 0{
		menor := vetAbs[0]
		for i, w:= range vetAbs{
			if w <= menor{
				menor = w
				indice = i
			}
		}
	vetSort = append(vetSort, vet[indice])
	vet = append(vet[:indice], vet[indice+1:]...)
	vetAbs = append(vetAbs[:indice], vetAbs[indice+1:]...)
	}
	return vetSort
}

func reverse(vet []int) []int {
	var vetReverse []int
	for i:=len(vet)-1; i>=0; i--{
		vetReverse = append(vetReverse, vet[i])
	}
	return vetReverse
}

func unique(vet []int) []int {
	qtd := len(vet)
	for i := 0; i < qtd; i++{
		count := 0
		for j := 0; j < qtd; j++{
			if vet[i] == vet[j]{
				count++
				if count > 1{
					vet = append(vet[:j], vet[j+1:]...)
					qtd--
					j--
				}
			}
		}
	}
	return vet
}

func repeated(vet []int) []int {
	var vet2 []int
	var count int
	qtd := len(vet)
	for i := 0; i< qtd; i++{
		menor := vet[0]
		for k :=0; k < qtd; k++{
			if vet[k] <= menor{
				menor = vet[k]
			}
		}
		// [5 3 3]
		// [1 3]
		count = 0
		for j := qtd-1; j>=0; j--{
			if vet[j] == menor{
				count++
				if count >= 2{
					vet2 = append(vet2, vet[j])
				}
				vet = append(vet[:j], vet[j+1:]...)
				qtd--
			}
		}
	}

	return vet2
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

