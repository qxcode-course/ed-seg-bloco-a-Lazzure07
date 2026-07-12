package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (node *Node) Next() *Node {
	if node.next == node.root {
		return nil
	}
	return node.next
}
func (node *Node) Prev() *Node {
	if node.prev == node.root {
		return nil
	}
	return node.prev
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
func (ll *Llist) Size() int {
	return ll.size
}
func (ll *Llist) PushBack(value int) {
	novo := &Node{
		value: value,
		next:  ll.root,
		prev:  ll.root.prev,
		root:  ll.root,
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
		root:  ll.root,
	}

	ll.root.next.prev = novo
	ll.root.next = novo

	ll.size++
}
func (ll *Llist) Clear() {
	ll.root.next = ll.root
	ll.root.prev = ll.root

	ll.size = 0
}

type Llist struct {
	root *Node
	size int
}

func (ll *Llist) Remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (ll *Llist) Insert(node *Node, newvalue int) {
	novo := &Node{
		value: newvalue,
		root:  ll.root,
		prev:  node.prev,
		next:  node,
	}

	node.prev.next = novo
	node.prev = novo
	ll.size++
}

func (ll *Llist) Search(oldvalue int) *Node {
	for node := ll.Front(); node != nil; node = node.Next() {
		if oldvalue == node.value {
			return node
		}
	}
	return nil
}

func (ll *Llist) Back() *Node {
	if ll.root.next == ll.root {
		return nil
	}
	return ll.root.prev
}

func (ll *Llist) Front() *Node {
	if ll.root.next == ll.root {
		return nil
	}
	return ll.root.next
}

type Node struct {
	value int
	next  *Node
	prev  *Node
	root  *Node
}

func NewLlist() *Llist {
	root := &Node{}
	root.next = root
	root.prev = root
	root.root = root

	return &Llist{
		root: root,
		size: 0,
	}
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLlist()

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
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
