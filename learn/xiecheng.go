package learn

import "fmt"

func Task(n int, ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("task-", n, ":", i)
	}
	ch <- n
}
