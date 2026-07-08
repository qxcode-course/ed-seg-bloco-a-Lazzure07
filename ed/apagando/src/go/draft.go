package main
import "fmt"
func main() {
	var numPessoas, nSaiu int
	var fila []int

	fmt.Scan(&numPessoas)
	for i:=0;i<numPessoas;i++{
		var id int
		fmt.Scan(&id)
		fila = append(fila, id)
	}
	fmt.Scan(&nSaiu)
	filaSaiu := make(map[int]bool, nSaiu)
	for i:=0;i<nSaiu;i++{
		var id int
		fmt.Scan(&id)
		filaSaiu[id] = true
	}
	
	for _, v := range fila{
		if filaSaiu[v] {
			continue
		}
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n")
}
