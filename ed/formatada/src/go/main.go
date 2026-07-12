package main

import (
	"bufio"
	"fmt"
	"os"

	"strconv"
	"strings"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// MyShow imprime a árvore binária de forma formatada.
func MyShow(node *Node, nivel int) {
	if node == nil{
		return
	}

	if node.Left != nil{
		MyShow(node.Left, nivel+1)
	}else if node.Right != nil{
		for range nivel+1{
			fmt.Print("....")
		}
		fmt.Println("#")
	}

	for range nivel{
		fmt.Print("....")

	}
	fmt.Println(strconv.Itoa(node.Value))

	if node.Right != nil{
		MyShow(node.Right, nivel+1)
	}else if node.Left!=nil{
		for range nivel + 1{
			fmt.Print("....")
		}
		fmt.Println("#")
	}
}
// -----------------------------------------------------------------------------------
func BShow(node *Node, history string) {
	if node != nil && (node.Left != nil || node.Right != nil) {
		BShow(node.Left, history+"l")
	}
	for i := 0; i < len(history)-1; i++ {
		if history[i] != history[i+1] {
			fmt.Print("│   ")
		} else {
			fmt.Print("    ")
		}
	}
	if history != "" {
		if history[len(history)-1] == 'l' {
			fmt.Print("╭───")
		} else {
			fmt.Print("╰───")
		}
	}
	if node == nil {
		fmt.Println("#")
		return
	}
	fmt.Println(node.Value)
	if node.Left != nil || node.Right != nil {
		BShow(node.Right, history+"r")
	}
}

func create(parts *[]string) *Node {
	// ["4", "#", "8", "2", "#", "#", "#"]
	if (*parts)[0] == "#" {
		*parts = (*parts)[1:]
		return nil
	}

	elem1, _ := strconv.Atoi((*parts)[0])
	*parts = (*parts)[1:]

	node := &Node{
		Value: elem1,		
	}

	node.Left   =  create(parts)
	node.Right  =  create(parts)

	return node
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	root := create(&parts)
	BShow(root, "") // Chama a função de impressão formatada
	MyShow(root, 0) // Chama a função de impressão personalizada
}
