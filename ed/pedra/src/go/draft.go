package main
import "fmt"
import "math"
import "slices"
func main() {
    var nParticipantes int
    var pedraA, pedraB float64
    var slice []float64

    fmt.Scan(&nParticipantes)

    for i := 1; i <= nParticipantes; i++{
        fmt.Scan(&pedraA, &pedraB)
        if pedraA < 10 || pedraB < 10{
            slice = append(slice, 100)
            continue
        }
        slice = append(slice, math.Abs(pedraA - pedraB))
    }
    min := slices.Min(slice)
    if min == 100{
        fmt.Println("sem ganhador")
    }else{
        for i, v := range slice{
            if min == v{
                fmt.Println(i)
            }
        }
    }
}
