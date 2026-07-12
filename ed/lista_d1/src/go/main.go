package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Llist struct {
	root *Node
	size int
}

func (ll *Llist) PopFront() {
	ll.root.next = ll.root.next.next
	ll.root.next.prev = ll.root

	ll.size--
}

func (ll *Llist) PopBack() {
	ll.root.prev = ll.root.prev.prev
	ll.root.prev.next = ll.root

	ll.size--
}

func (ll *Llist) Clear() {
	ll.root.next = ll.root
	ll.root.prev = ll.root

	ll.size = 0
}

func (ll *Llist) Size() int {
	return ll.size
}
func (ll *Llist) PushBack(value int) {
	novo := &Node{
		value: value,
		next:  ll.root,
		prev:  ll.root.prev,
	}

	ll.root.prev.next = novo
	ll.root.prev = novo

	ll.size++
}
func (ll *Llist) PushFront(value int) {
	novo := &Node{
		value: value,
		next:  ll.root.next,
		prev:  ll.root,
	}

	ll.root.next.prev = novo
	ll.root.next = novo

	ll.size++
}
func (ll *Llist) String() string {
	s := ""

	node := ll.root.next
	for node != ll.root {
		s += strconv.Itoa(node.value)

		if node.next != ll.root {
			s += ", "
		}

		node = node.next
	}

	return "[" + s + "]"
}

type Node struct {
	value int
	next  *Node
	prev  *Node
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

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
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}

func NewLList() *Llist {
	root := &Node{}
	root.next = root
	root.prev = root

	return &Llist{
		root: root,
		size: 0,
	}
}
