package main
import "fmt"
func main() {
    var heli, police, fug, dir int

    fmt.Scan(&heli, &police, &fug, &dir)

    for{
        if fug == 16{
            fug = 0
        }
        if fug == -1{
            fug = 15
        }
        if dir == 1{
            if fug == police{
                fmt.Println("N")
                break
            }
            if fug == heli{
                fmt.Println("S")
                break
            }
            fug++
        }else{
            if fug == police{
                fmt.Println("N")
                break
            }
            if fug == heli{
                fmt.Println("S")
                break
            }
            fug--
        }
    }


}
