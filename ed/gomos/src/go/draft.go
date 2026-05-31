package main

import "fmt"
func main() {
    var intQ, posX, posY int
    var charD string
    var matCobra [][]int

    fmt.Scan(&intQ, &charD)
    for i := 0; i<intQ; i++{
        fmt.Scan(&posX, &posY)
        if i == 0{
            switch(charD){
                case "U":
                    matCobra = append(matCobra, []int{posX, posY-1})
                case "D":
                    matCobra = append(matCobra, []int{posX, posY+1})
                case "R":
                    matCobra = append(matCobra, []int{posX+1, posY})
                case "L":
                    matCobra = append(matCobra, []int{posX-1, posY})
            }
        }
        matCobra = append(matCobra, []int{posX, posY})
    }

    for i, linha := range matCobra{
        if i == len(matCobra)-1{
            break
        }
        for j, v:= range linha{
            if j==0{
                fmt.Printf("%d ", v)
            }else{
                fmt.Printf("%d\n", v)
            }
        }
    }
}
