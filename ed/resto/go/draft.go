package main
import "fmt"

func binario(num int){
    if num == 0{
        return
    }
    resto := num%2
    resultado := num/2

    binario(resultado)

    fmt.Println(resultado, resto)
}

func main() {
    var num int
    fmt.Scan(&num)
    binario(num)
}
