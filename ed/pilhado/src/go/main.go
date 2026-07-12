package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
	dir  int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var linhas, colunas int
	fmt.Fscan(in, &linhas, &colunas)
	in.ReadByte()

	lab := make([][]byte, linhas)

	var inicio, fim Pos

	for i := 0; i < linhas; i++ {
		linha, _ := in.ReadBytes('\n')
		if len(linha) > 0 && linha[len(linha)-1] == '\n' {
			linha = linha[:len(linha)-1]
		}
		lab[i] = linha

		for j := 0; j < colunas; j++ {
			if lab[i][j] == 'I' {
				inicio = Pos{l: i, c: j}
			}
			if lab[i][j] == 'F' {
				fim = Pos{l: i, c: j}
			}
		}
	}
	lab[inicio.l][inicio.c] = ' '
	lab[fim.l][fim.c] = ' '

	visitado := make([][]bool, linhas)
	for i := 0; i < linhas; i++ {
		visitado[i] = make([]bool, colunas)
	}

	mov := []struct{ dl, dc int }{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	caminho := NewStack[Pos]()
	visitado[inicio.l][inicio.c] = true
	caminho.Push(inicio)

	for !caminho.IsEmpty() {
		atual := caminho.Pop()

		if atual.l == fim.l && atual.c == fim.c {
			caminho.Push(atual)
			break
		}

		if atual.dir == 4 {
			continue
		}

		prox := atual
		prox.dir++
		caminho.Push(prox)

		nl := atual.l + mov[atual.dir].dl
		nc := atual.c + mov[atual.dir].dc

		if nl >= 0 && nl < linhas &&
			nc >= 0 && nc < colunas &&
			lab[nl][nc] != '#' &&
			!visitado[nl][nc] {

			visitado[nl][nc] = true
			caminho.Push(Pos{l: nl, c: nc})
		}
	}

	for !caminho.IsEmpty() {
		p := caminho.Pop()
		if lab[p.l][p.c] == ' ' {
			lab[p.l][p.c] = '.'
		}
	}

	for i := 0; i < linhas; i++ {
		fmt.Println(string(lab[i]))
	}
}