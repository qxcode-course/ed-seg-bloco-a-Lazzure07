package main
import (
    "fmt"
    //"math"
)
func matar(e int, sliceFila []int)(int, []int){
    for i:=0;i<len(sliceFila);i++{
        if e == i{
            if i == len(sliceFila)-1{
                sliceFila = append(sliceFila[:0], sliceFila[1:]...)
            }else{
                sliceFila = append(sliceFila[:i+1], sliceFila[i+2:]...)
            }
            break
        }
    }
    e++
    if e >= len(sliceFila){ e = 0 }
    return e, sliceFila
}
func toString(e int, sliceFila []int){
    fmt.Printf("[ ")
    for i:= 0; i<len(sliceFila);i++{
        if i == e{
            fmt.Printf("%d> ", sliceFila[i])
            continue
        }
        fmt.Printf("%d ", sliceFila[i])
    }
    fmt.Printf("]\n")
}
func main() {
    var n, e int
    var sliceFila[] int

    fmt.Scan(&n, &e)
    e--
    // ABAIXO É O CODIGO QUE FIZ PARA CALCULAR APENAS O GANHADOR
 
    // resultado := (((2*n - int((math.Pow(2, float64(int(math.Log2(float64(n))) + 1)))) + e) - 1) % n) + 1
    // fmt.Println(resultado)

    for i := 1; i<=n;i++{
        sliceFila = append(sliceFila, i)
    }
    for i := 0; i<n;i++{
        toString(e, sliceFila)
        e, sliceFila = matar(e, sliceFila)
    }
}