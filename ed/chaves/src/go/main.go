package main

import (
	"fmt"
)

func main() {
	q := NewQueue[string]()

	for i := 'A'; i <= 'P'; i++{
		q.Enqueue(string(i))
	}

	for range 15 {
		time1 := q.Dequeue()
		time2 := q.Dequeue()
		gols1 := 0
		gols2 := 0

		fmt.Scan(&gols1, &gols2)
		if gols1>gols2{
			q.Enqueue(time1)
		}else{
			q.Enqueue(time2)
		}
	}
	fmt.Println(q.Dequeue())
}
