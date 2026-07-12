package main
import "fmt"
func main() {
    var seq string
    var l int
    fmt.Scan(&seq, &l)

    vetor := []byte(seq)

    backtrack(vetor, 0, l)

    fmt.Println(string(vetor))
}

func backtrack (vet []byte, pos, l int)bool{
    if pos == len(vet){
        return true
    }

    if vet[pos] != '.'{
        return backtrack(vet, pos+1, l)
    }

    for i := byte('0'); i <= byte('0'+l); i++ {
		if pode(vet, pos, i, l) {
			vet[pos] = i

			if backtrack(vet, pos+1, l) {
				return true
			}

			vet[pos] = '.'
		}
    }
    return false
}

func pode(vet []byte, pos int, v byte, l int) bool {
	inicio := pos - l
    fim := pos + l
    if inicio < 0{inicio = 0}
    if fim > len(vet) - 1 {fim = len(vet) -1}

    for i := inicio; i <= fim; i++{
        if vet[i] == v && i != pos{
            return false
        }
    }
    return true
}