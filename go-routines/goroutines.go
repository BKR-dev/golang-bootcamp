package goroutines

import "fmt"

func sum(a, b int, c chan int) {
	c <- a + b
}

func GoRoutinesExample1() {
	c := make(chan int)

	go sum(1, 2, c)
	go sum(3, 4, c)
	go sum(2, 3, c)

	x, y, z := <-c, <-c, <-c

	fmt.Println("Results: ", x, y, z)
}
