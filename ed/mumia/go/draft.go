package main

import "fmt"
func main() {
    var nome string
    var idade int

    fmt.Scan(&nome, &idade)

    if idade < 12{
        fmt.Printf("%s eh crianca\n", nome)
    }
    if idade >= 12 && idade <18{
        fmt.Printf("%s eh jovem\n", nome)
    }
    if idade >= 18 && idade <65{
        fmt.Printf("%s eh adulto\n", nome)
    }
    if idade >= 65 && idade <1000{
        fmt.Printf("%s eh idoso\n", nome)
    }
    if idade >= 1000{
        fmt.Printf("%s eh mumia\n", nome)
    }
}
