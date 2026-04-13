package main
import "fmt"
func main() {
    var numN, numM, id , idQuit int
    var fila []int

    fmt.Scan(&numN)
    for i:=0; i<numN; i++{
        fmt.Scan(&id)
        fila = append(fila, id)
    }
	fmt.Scan(&numM)
    for i:=0; i<numM; i++{
        fmt.Scan(&idQuit)
        for i:=0; i<len(fila); i++{
			if fila[i] == idQuit{
                fila = append(fila[:i], fila[i+1:]...)
				i--
			}
		}
    }

	for _, v := range fila{
	    fmt.Printf("%d ", v)
	}
	fmt.Printf("\n")


}
