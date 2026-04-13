package main
import "fmt"
func main() {
    var num, animal, conta int
    var sliceAnimais []int

    fmt.Scan(&num)
    for i:=0; i<num; i++{
        fmt.Scan(&animal)
        sliceAnimais = append(sliceAnimais, animal)
    }

    for i := 0; i < len(sliceAnimais); i++{
        for j := i+1; j < len(sliceAnimais); j++{
            if sliceAnimais[j] == sliceAnimais[i]*(-1){
                sliceAnimais = append(sliceAnimais[:j], sliceAnimais[j+1:]...)
                sliceAnimais = append(sliceAnimais[:i], sliceAnimais[i+1:]...)
                conta++
                i--
                break
            }
        }
    }
    fmt.Println(conta)

}
