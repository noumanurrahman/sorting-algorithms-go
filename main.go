package main

import (
	"fmt"
	"sorting-algorithms/algorithms"
)

func main() {
	nums := []int{5, 2, 9, 4, 6, 3}
	algorithms.BubbleSort(nums)
	fmt.Println("Bubble Sort:", nums)
}
