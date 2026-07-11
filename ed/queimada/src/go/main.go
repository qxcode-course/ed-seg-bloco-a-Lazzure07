package main

import (
	"bufio"
	"fmt"
	"os"
)

func burnTrees(grid [][]rune, lfire, cfire int) {
	nl := len(grid)
	nc := len(grid[0])
	// _, _ = nl, nc
	// se estiver fora da matriz, retorne
	// se o elemento atual não for uma arvore, retorne
	if lfire < 0 || lfire >= nl || cfire < 0 || cfire >= nc || grid[lfire][cfire] != '#'{
		return
	}
	// queime a arvore colocando o caractere 'o' na posição atual
	grid[lfire][cfire] = 'o'
	// chame a recursão para todos os 4 vizinhos
	burnTrees(grid, lfire -1, cfire)
	burnTrees(grid, lfire, cfire-1)
	burnTrees(grid, lfire +1, cfire)
	burnTrees(grid, lfire , cfire+1)

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(grid [][]rune) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
