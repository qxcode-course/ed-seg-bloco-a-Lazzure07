package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

// func getNeig(p Pos) []Pos {
// 	return []Pos{{p.l, p.c - 1}, {p.l - 1, p.c}, {p.l, p.c + 1}, {p.l + 1, p.c}}
// }

// func inside(grid [][]rune, p Pos) bool {
// 	return !(p.l < 0 || p.l >= len(grid) || p.c < 0 || p.c >= len(grid[0]))
// }

// func match(grid [][]rune, p Pos, value rune) bool {
// 	return inside(grid, p) && grid[p.l][p.c] == value
// }

// Função recursiva que tenta encontrar o caminho do início ao fim
func search(grid [][]rune, startPos, endPos Pos) bool {
	_, _, _ = grid, startPos, endPos

	visitado := make([][]bool, len(grid))
	for i := range grid{
		visitado[i] = make([]bool, len(grid[0]))
	}


	var procura func(grid [][]rune, linha, coluna int, final Pos) bool

	procura = func(grid [][]rune, linha, coluna int, final Pos)bool{

		if linha < 0 || linha >= len(grid) || coluna < 0 || coluna >= len(grid[0]) || grid[linha][coluna] == '#' || visitado[linha][coluna]{
			return false
		}

		if linha == final.l && coluna == final.c{
			// achei
			grid[linha][coluna] = '.'
			return true
		}

		visitado[linha][coluna] = true

		if procura(grid, linha-1, coluna, final) || procura(grid, linha+1, coluna, final) || procura(grid, linha, coluna-1, final) || procura(grid, linha, coluna+1, final){
			grid[linha][coluna] = '.'
			return true
		}

		return false
	}

	procura(grid, startPos.l, startPos.c, endPos)

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nl_nc := scanner.Text()
	var nl, nc int
	fmt.Sscanf(nl_nc, "%d %d", &nl, &nc)
	grid := make([][]rune, nl)

	// Lê a gridriz
	for i := range nl {
		scanner.Scan()
		grid[i] = []rune(scanner.Text())
	}

	// Procura posições de início e endPos e conserta para _
	var startPos, endPos Pos
	for l := range nl {
		for c := range nc {
			if grid[l][c] == 'I' {
				grid[l][c] = ' '
				startPos = Pos{l, c}
			}
			if grid[l][c] == 'F' {
				grid[l][c] = ' '
				endPos = Pos{l, c}
			}
		}
	}

	search(grid, startPos, endPos)

	// Imprime o labirinto final
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
