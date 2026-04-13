package main
import "fmt"
func main() {
    var num, numAnim, count int
    var Animais []int

    fmt.Scan(&num)
    for i:=0; i<num; i++{
        fmt.Scan(&numAnim)
        Animais = append(Animais, numAnim)
    }

    for i := 0; i < len(Animais); i++{
        for j := i+1; j < len(Animais); j++{
            if Animais[j] == Animais[i]*(-1){
                Animais = append(Animais[:j], Animais[j+1:]...)
                Animais = append(Animais[:i], Animais[i+1:]...)
                count++
                i--
                break
            }
        }
    }
    fmt.Println(count)

}
