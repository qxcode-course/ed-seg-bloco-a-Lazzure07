package main
import "fmt"
func backtracking (conjunto []int, value, index, sum int)bool{

    if sum == value{
        return true
    }
    
    if index == len(conjunto){
        return false
    }

	if backtracking(conjunto, value, index+1, sum + conjunto[index]) {
		return true
	}

    if backtracking(conjunto, value, index+1, sum){
        return true
    }

    return false
}
func main() {
    var size, value int
    fmt.Scan(&size, &value)
    conjunto := make([]int, size)
    for i := range size {
        fmt.Scan(&conjunto[i])
    }

    fmt.Println(backtracking(conjunto, value, 0, 0))

}