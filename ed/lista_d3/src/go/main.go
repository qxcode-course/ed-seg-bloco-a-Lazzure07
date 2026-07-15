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
	next  *Node
	prev  *Node
	root  *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	list := &LList{}
	list.root = &Node{root: nil}
	list.root.next = list.root
	list.root.prev = list.root
	list.root.root = list.root // nó sentinela aponta pra si mesmo
	return list
}

func (l *LList) PushBack(value int) {
	l.insertBefore(l.root, value)
}

func (l *LList) insertBefore(mark *Node, value int) {
	n := &Node{
		Value: value,
		root:  l.root,
	}
	n.prev = mark.prev
	n.next = mark
	mark.prev.next = n
	mark.prev = n
}


func str2list(serial string) *LList {
	serial = serial[1 : len(serial)-1]
	ll := NewLList()
	if serial == "" {
		return ll
	}
	for _, p := range strings.Split(serial, ",") {
		value, _ := strconv.Atoi(p)
		ll.PushBack(value)
	}
	return ll
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "compare":
			lla := str2list(args[1])
			llb := str2list(args[2])
			if equals(lla, llb) {
				fmt.Println("iguais")
			} else {
				fmt.Println("diferentes")
			}
		case "addsorted":
			lla := NewLList()
			for i := 1; i < len(args); i++ {
				value, _ := strconv.Atoi(args[i])
				addsorted(lla, value)
			}
			fmt.Println(toString(lla))
		case "reverse":
			lla := str2list(args[1])
			fmt.Println(toString(reverse(lla)))
		case "merge":
			lla := str2list(args[1])
			llb := str2list(args[2])
			merged := merge(lla, llb)
			fmt.Println(merged)
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}

func reverse(lla *LList)*LList {
	llv := NewLList()

	for node := lla.root.prev; node != lla.root; node = node.prev {
		llv.PushBack(node.Value)
	}

	return llv
}

func addsorted(lla *LList, value int) {
	lastValue := lla.root.prev
	
	if lastValue == lla.root{
		lla.PushBack(value)
		return
	}

	var add func(lastvalue *Node)

	add = func(lastvalue *Node){
		if lastvalue == lla.root{
			lla.insertBefore(lastvalue.next, value)
			return
		}
		if value < lastvalue.Value{
			add(lastvalue.prev)
			return
		}
		if value >= lastvalue.Value{
			lla.insertBefore(lastvalue.next, value)
			return
		}
	}

	add(lastValue)
}

func toString(llv *LList)string{
	if llv.root.next == llv.root{
		return "[]"
	}

	s := "["
	for i := llv.root.next; i != llv.root; i = i.next {
		s = s + strconv.Itoa(i.Value) + ", "
	}
	s = s[:len(s)-2]
	return s + "]"
}

func merge(lla, llb *LList) any {

	llv := "["

	var compara func(nodeA, nodeB *Node)

	compara = func(nodeA, nodeB *Node){

		if nodeA == lla.root && nodeB == llb.root{
			return
		}
		if (nodeA != lla.root && nodeB == llb.root) {
			llv += strconv.Itoa(nodeA.Value) +", "
			compara(nodeA.next, nodeB)
			return
		}
		if (nodeA == lla.root && nodeB != llb.root) {
			llv += strconv.Itoa(nodeB.Value)+", "
			compara(nodeA, nodeB.next)
			return
		}
		if nodeA.Value < nodeB.Value{
			llv += strconv.Itoa(nodeA.Value)+", "
			compara(nodeA.next, nodeB)
			return
		}
		if nodeB.Value < nodeA.Value{
			llv += strconv.Itoa(nodeB.Value)+", "
			compara(nodeA, nodeB.next)
			return
		}
		if nodeA.Value == nodeB.Value{
			llv = llv + strconv.Itoa(nodeA.Value)+", " + strconv.Itoa(nodeA.Value)+", "
			compara(nodeA.next, nodeB.next)
			return
		}

	}
	compara(lla.root.next, llb.root.next)

	llv = llv[:len(llv)-2]
	return  llv +"]"
}

func equals(lla, llb *LList) bool {

	var compara func(nodeA, nodeB *Node)bool

	compara = func(nodeA, nodeB *Node)bool{

		if nodeA == lla.root && nodeB == llb.root {
			return true
		}
		if nodeA == nil || nodeB == nil{
			return false
		}
		if nodeA.Value != nodeB.Value{
			return false
		}
		return compara(nodeA.next, nodeB.next)
	}

	return compara(lla.root.next, llb.root.next)
}
