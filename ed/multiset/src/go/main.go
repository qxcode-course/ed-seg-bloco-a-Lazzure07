package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MultiSet struct {
	data     []int
	size     int
	capacity int
}

func (ms *MultiSet) Clear() {
	ms.size = 0
	ms.data = nil
}

func (ms *MultiSet) Unique() int {
	count := 0
	if ms.size == 0 {
		return count
	}
	for i := 1; i < ms.size; i++ {
		if ms.data[i] != ms.data[i-1] {
			count++
		}
	}
	return count + 1
}

func (ms *MultiSet) Count(value int) int {
	index, existe := ms.binarySearch(value)
	if !existe {
		return 0
	}

	count := 1

	for i := index - 1; i >= 0 && ms.data[i] == value; i-- {
		count++
	}

	for i := index + 1; i < ms.size && ms.data[i] == value; i++ {
		count++
	}

	return count
}

func (ms *MultiSet) Erase(value int) {
	index, existe := ms.binarySearch(value)
	if !existe {
		fmt.Println("value not found")
		return
	}

	ms.data = append(ms.data[:index], ms.data[index+1:]...)
	ms.size--
}

func (ms *MultiSet) Insert(value int) {

	index, _ := ms.binarySearch(value)

	ms.data = append(ms.data[:index], append([]int{value}, ms.data[index:]...)...)
	ms.size++

}
func (ms *MultiSet) binarySearch(value int) (int, bool) {
	inicio := 0
	fim := ms.size - 1

	for inicio <= fim {
		meio := (inicio + fim) / 2

		if ms.data[meio] == value {
			return meio, true
		} else if inicio == fim && value < ms.data[meio] {
			return inicio, false
		} else if inicio == fim && value > ms.data[meio] {
			return inicio + 1, false
		}

		if value < ms.data[meio] {
			fim = meio - 1
		} else {
			inicio = meio + 1
		}
	}

	return 0, false
}
func (ms *MultiSet) Contains(value int) bool {
	for _, num := range ms.data {
		if num == value {
			return true
		}
	}
	return false
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

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(Join(ms.data, ", "))
		case "erase":
			value, _ := strconv.Atoi(args[1])
			ms.Erase(value)
		case "contains":
			value, _ := strconv.Atoi(args[1])
			if ms.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}

func NewMultiSet(i int) *MultiSet {
	return &MultiSet{
		data:     make([]int, 0, i),
		size:     0,
		capacity: i,
	}
}
