package main
import (
    "fmt"
    //"math"
)
func matar(e int, sliceFila []int)(int, []int){
    for i, v := range sliceFila{
        if e == i && v > 0{
            if i == len(sliceFila)-1{
                sliceFila = append(sliceFila[:0], sliceFila[1:]...)
            }else{
                sliceFila = append(sliceFila[:i+1], sliceFila[i+2:]...)
            }
            e++
            break
        }else if e == i && v < 0{
            if i == 0{
                sliceFila = sliceFila[:len(sliceFila)-1]
            }else{
                sliceFila = append(sliceFila[:i-1], sliceFila[i:]...)
            }
            e -= 2
            break
        }
    }
    if e >= len(sliceFila){ e = 0 }
    if e < 0{ e = len(sliceFila)-1 }
    return e, sliceFila
}
func toString(e int, sliceFila []int){
    fmt.Printf("[ ")
    for i:= 0; i<len(sliceFila);i++{
        if (i == e || i == -e) && sliceFila[i] > 0{
            fmt.Printf("%d> ", sliceFila[i])
            continue
        }else if (i == e || i == -e) && sliceFila[i] < 0{
            fmt.Printf("<%d ", sliceFila[i])
            continue
        }
        fmt.Printf("%d ", sliceFila[i])
    }
    fmt.Printf("]\n")
}
func main() {
    var n, e, f int
    var sliceFila[] int

    fmt.Scan(&n, &e, &f)
    e--
    // ABAIXO É O CODIGO QUE FIZ PARA CALCULAR APENAS O GANHADOR
 
    // resultado := (((2*n - int((math.Pow(2, float64(int(math.Log2(float64(n))) + 1)))) + e) - 1) % n) + 1
    // fmt.Println(resultado)

    for i := range n{
        sliceFila = append(sliceFila, (i+1)*f)
        f *= -1
    }

    for len(sliceFila) > 0{
        toString(e, sliceFila)
        e, sliceFila = matar(e, sliceFila)
    }
}