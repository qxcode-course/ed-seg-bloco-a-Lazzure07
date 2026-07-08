package main
import "fmt"

func inverterVetor(vetor[]int, inicio, fim int){
    if inicio >= fim{
        return
    }
    vetor[inicio], vetor[fim] = vetor[fim], vetor[inicio]
    inicio++
    fim--

    inverterVetor(vetor, inicio, fim)
}

func main() {
    var n, rota int
    fmt.Scan(&n, &rota)
    rota %= n

    vetor := make([]int, n)
    for i:= range n{
        fmt.Scan(&vetor[i])
    }

    inverterVetor(vetor, 0, len(vetor)-1)
    inverterVetor(vetor, 0, rota-1)
    inverterVetor(vetor, rota, len(vetor)-1)

    fmt.Printf("[ ")
    for _, v := range vetor{
        fmt.Printf("%d ", v)
    }
    fmt.Println("]")
}
