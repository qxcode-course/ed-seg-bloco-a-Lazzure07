package main
import "fmt"
func main() {
    var numFig, numAlb, fig int
    var sliceBaruel []int
    var sliceRepet []int
    var sliceFalta []int
    var verificador bool


    fmt.Scan(&numAlb, &numFig)
    for i:= 0; i < numFig; i++{
        fmt.Scan(&fig)
        sliceBaruel = append(sliceBaruel, fig)
    }

    for i:= 0; i<len(sliceBaruel); i++{
        for j := i +1; j<len(sliceBaruel); j++{
            if(sliceBaruel[i] == sliceBaruel[j]){
                sliceRepet = append(sliceRepet, sliceBaruel[j])
                sliceBaruel = append(sliceBaruel[:j], sliceBaruel[j+1:]...)
                j--
            }
        }
    }

    for j := 1; j<=numAlb; j++{
        verificador = false
        for  i := 0; i<len(sliceBaruel); i++{

            if sliceBaruel[i] == j{
                verificador = true
                break
            }
        }
        if !verificador{
            sliceFalta = append(sliceFalta, j)
        }
    }

    if sliceRepet == nil{
        fmt.Println("N")
    }else{
        for i, v := range sliceRepet{
            if i == len(sliceRepet)-1{
                fmt.Printf("%d", v)
                break
            } 
            fmt.Printf("%d ", v)
        }
        fmt.Printf("\n")
    }

    if sliceFalta == nil{
        fmt.Println("N")
    }else{
        for i, v := range sliceFalta{
            if i == len(sliceFalta)-1{
                fmt.Printf("%d", v)
                break
            } 
            fmt.Printf("%d ", v)
        }
        fmt.Printf("\n")

    }

}
