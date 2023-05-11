package main

import (
	"fmt"

	gd "github.com/BKR-dev/golang-bootcamp/go-datastructures"
)

func main() {
	// gr.GoRoutinesExample1()
	// gr.GoRoutinesExample2()
	var nums = []int{1, 2, 3, 4}
	var target = 6

	fmt.Println(gd.TwoSum(nums, target))
}
