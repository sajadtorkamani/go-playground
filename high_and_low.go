package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func HighAndLow(str string) string {
	chars := strings.Split(str, " ")
	var nums []int

	for _, char := range chars {
		num, err := strconv.Atoi(char)

		if err != nil {
			panic(err)
		}

		nums = append(nums, num)
	}

	sort.Ints(nums)

	return fmt.Sprintf("%d %d", nums[len(nums)-1], nums[0])
}

func main() {
	fmt.Println(HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4"))
}
