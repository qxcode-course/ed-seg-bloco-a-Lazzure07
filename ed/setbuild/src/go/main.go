package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data     []int
	size     int
	capacity int
}

func (s *Set) Erase(value int) bool {
	index, existe := s.binarySearch(value)
	if !existe {
		return false
	}

	s.data = append(s.data[:index], s.data[index+1:]...)
	s.size--
	return true
}

func (s *Set) Contains(value int) bool {
	for _, num := range s.data {
		if num == value {
			return true
		}
	}
	return false
}
func (s *Set) binarySearch(value int) (int, bool) {
	inicio := 0
	fim := s.size - 1

	for inicio <= fim {
		meio := (inicio + fim) / 2

		if s.data[meio] == value {
			return meio, true
		} else if inicio == fim && value < s.data[meio] {
			return inicio, false
		} else if inicio == fim && value > s.data[meio] {
			return inicio + 1, false
		}

		if value < s.data[meio] {
			fim = meio - 1
		} else {
			inicio = meio + 1
		}
	}

	return -1, false
}
func (s *Set) Insert(value int) {

	index, existe := s.binarySearch(value)
	if existe {
		return
	}
	if index == -1 {
		index = 0
	}
	s.data = append(s.data[:index], append([]int{value}, s.data[index:]...)...)
	s.size++

}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			fmt.Println(Join(v.data, ", "))
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			err := v.Erase(value)
			if !err{
				fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}

func NewSet(capacity int) *Set {
	return &Set{
		data:     make([]int, 0, capacity),
		size:     0,
		capacity: capacity,
	}
}
func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return "[]"
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return "[" + result.String() + "]"
}
