package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {

	// [["A","B","C","E"],
	//  ["S","F","C","S"],
	//  ["A","D","E","E"]], word = "SEE"

	visitado := make([][]bool, len(grid))
	for i := range visitado{
		visitado[i] = make([]bool, len(grid[0]))
	}

	var backtracking func(linha, coluna, indice int)bool

	backtracking = func(linha, coluna, indice int)bool{

		if indice == len(word){
			return true
		}

		if linha < 0  || linha >= len(grid)     || 
		   coluna < 0 || coluna >= len(grid[0]) || 
		   visitado[linha][coluna]              || 
		   grid[linha][coluna] != word[indice]   { return false }

		visitado[linha][coluna] = true

		if backtracking(linha-1, coluna, indice+1) || 
		   backtracking(linha+1, coluna, indice+1) || 
		   backtracking(linha, coluna-1, indice+1) || 
		   backtracking(linha, coluna+1, indice+1){
			return true
		   }

		   visitado[linha][coluna] = false

		return false
	}

	for i, line := range grid{
		for j := range line{
			if backtracking(i, j, 0){
				return true
			}
		}
	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
