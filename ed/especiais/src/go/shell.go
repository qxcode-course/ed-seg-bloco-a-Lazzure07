package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	One int
	Two int
}

func occurr(vet []int) []Pair {
	var count int
	var result []Pair
	var vetAbs []int
	for _, v := range vet{
		if v < 0{
			vetAbs = append(vetAbs, v*(-1))
		}else{
			vetAbs = append(vetAbs, v)
		}
	}
	qtd:=len(vet)
	i := 0
	for i < qtd{
		menor := vetAbs[0]
		for m := qtd-1; m >= 0; m--{
			if menor > vetAbs[m]{
				menor = vetAbs[m]
			}
		}
		for j := qtd-1; j >= 0; j--{
			if menor == vetAbs[j]{
				count++
			}
		}
		result = append(result, Pair{One: menor, Two: count})
		
		for k := qtd-1; k >= 0; k--{
			if menor == vetAbs[k]{
				vetAbs = append(vetAbs[:k], vetAbs[k+1:]...)
			}
		}
		qtd -= count
		count = 0
	}
	return result
}

func teams(vet []int) []Pair {
	var result []Pair
	var count, j int
	if vet == nil{
		return result
	}else{
		for i:=0; i < len(vet); i+=count{
			count = 0
			for j = i; j < len(vet); j++{
				if vet[i] == vet[j]{
					count++
					if j == len(vet)-1 {
						result = append(result, Pair{One: vet[i], Two: count})
						break
					}
				}else{
					result = append(result, Pair{One: vet[i], Two: count})
					break
				}
			}
		}
		return result
	}
}

func mnext(vet []int) []int {
	qtd := len(vet)
	if qtd == 0 {
		return vet
	}
	for i := 0; i < qtd; i++{
		if vet[i] > 0 && qtd > 1{ 
			if i > 0 && i < qtd-1{
				if vet[i-1] < 0 || vet[i+1] < 0{
					vet[i] = 1
				}else{
					vet[i] = 0
				}
			}else if i == 0{
				if vet[i+1] < 0{
					vet[i] = 1
				}else{
					vet[i] = 0
				}
			}else{
				if vet[i-1] < 0{
					vet[i] = 1
				}else{
					vet[i] = 0
				}
			}
		}else if qtd == 1{
			vet[0] = 0
		}
	}
	for j := 0; j < qtd; j++{
		if vet[j] < 0{
			vet[j] = 0
		}
	}
	return vet
}

func alone(vet []int) []int {
	qtd := len(vet)
	if qtd == 0 {
		return vet
	}
	for i := 0; i < qtd; i++{
		if vet[i] > 0 && qtd > 1{ 
			if i > 0 && i < qtd-1{
				if vet[i-1] < 0 || vet[i+1] < 0{
					vet[i] = 0
				}else{
					vet[i] = 1
				}
			}else if i == 0{
				if vet[i+1] < 0{
					vet[i] = 0
				}else{
					vet[i] = 1
				}
			}else{
				if vet[i-1] < 0{
					vet[i] = 0
				}else{
					vet[i] = 1
				}
			}
		}else if qtd == 1 && vet[i] > 0{
			vet[0] = 1
		}
	}
	for j := 0; j < qtd; j++{
		if vet[j] < 0{
			vet[j] = 0
		}
	}
	return vet
}

func couple(vet []int) int {
	var count int
	// [0,0,4,2,-4,-4,4]
	for i := 0; i < len(vet); i++{
		for j := 0; j < len(vet); j++{
			if vet[j] == vet[i] * (-1) && vet[i] != 0{
				count++
				vet[i] = 0
				vet[j] = 0
				break
			}
		}
	}
	return count
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	_ = vet
	_ = seq
	_ = pos
	return false
}

func subseq(vet []int, seq []int) int {
	var count, indiceInicio, indice int = 0, -1, -1

	for i := 0; i < len(seq); i++{
		for j := indice+1; j < len(vet); j++{
			if seq[i] == vet[j]{
				count++
				indice = j
				if i == 0{
					indiceInicio = j
				}
				break
			}else if indiceInicio != -1{
				vet[indiceInicio] = -1
				i = -1
				count = 0
				indice = -1
				indiceInicio = -1
				break
			}
		}
		if count == len(seq){
			return indiceInicio
		}
	}
	return -1
}

func erase(vet []int, posList []int) []int {
	for i := 0; i < len(posList); i++{
		for j := 0; j < len(vet); j++{
			if posList[i] == j{
				vet[j] = -1
				break
			}
		}
	}
	qtd := len(vet)
	for k := 0; k < qtd; k++{
		if vet[k] == -1{
			vet = append(vet[:k], vet[k+1:]...)
			qtd--
			k--
		}
	}
	return vet
}

func clear(vet []int, value int) []int {

	qtd := len(vet)
	for k := 0; k < qtd; k++{
		if vet[k] == value{
			vet = append(vet[:k], vet[k+1:]...)
			qtd--
			k--
		}
	}
	return vet
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}