package main

import "fmt"

func PositivesSum(nums []int) int {
	sum := 0

	for _, num := range nums {
		if num > 0 {
			sum += num
		}
	}

	return sum
}

func main() {
	fmt.Println(PositivesSum([]int{2, 4, 6, -10, -20}))
}
