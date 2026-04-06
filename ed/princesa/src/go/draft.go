package main
import (
    "fmt"
    "math"
)
func main() {
    var n, x, y int

    fmt.Scan(&n, &x)
    varAux := n
    for {
        if(varAux/2==0){
            y++
            break
        }
        y++
        varAux = varAux/2
    }
    resultado := 2*n - int((math.Pow(2, float64(y)))) + x
    for resultado>n{
        resultado -= n
    }
    fmt.Println(resultado)
}