package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct{
	l int
	c int
}

func burnTrees(grid [][]rune, l, c int) {
	stack := NewStack[Pos]()
	stack.Push(Pos{l: l, c: c})

	for !stack.IsEmpty(){
		coord := stack.Pop()

		if coord.l < 0 || coord.l >= len(grid) ||
		   coord.c < 0 || coord.c >= len(grid[0])||
		   grid[coord.l][coord.c] != '#'{
			continue
		}

		grid[coord.l][coord.c] = 'o'

		stack.Push(Pos{ l: coord.l -1, c: coord.c})
		stack.Push(Pos{ l: coord.l +1, c: coord.c})
		stack.Push(Pos{ l: coord.l, c: coord.c -1})
		stack.Push(Pos{ l: coord.l, c: coord.c +1})

	}

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

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
