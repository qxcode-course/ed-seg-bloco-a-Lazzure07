package main
import "fmt"
func main() {

    var n int
    mapa := make(map[string]int)
    var input string
    var output []int

    fmt.Scan(&n)
    for range n{
        fmt.Scan(&input)
        mapa[input]++
    }

    fmt.Scan(&n)
    for range n{
        fmt.Scan(&input)
        output = append(output, mapa[input])
    }

    for i, v := range output{
        if i > 0{
            fmt.Print(" ")
        }
        fmt.Print(v)
    }
    fmt.Println()

}
