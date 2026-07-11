package main

import (
	"fmt"
	"sorting-algorithms/algorithms"
)

func main() {
	nums := []int{5, 2, 9, 4, 6, 3}
	algorithms.BubbleSort(nums)
	fmt.Println("Bubble Sort:", nums)
	algorithms.InsertionSort(nums)
	fmt.Println("Insertion Sort:", nums)
	algorithms.SelectionSort(nums)
	fmt.Println("Selection Sort:", nums)
	sortedNums := algorithms.MergeSort(nums)
	fmt.Println("Merge Sort:", sortedNums)

	fibExponential := algorithms.FibonacciExponential(10)
	fmt.Println("Fibonacci Exponential:", fibExponential)
	fibPolynomial := algorithms.Fibonacci(100)
	fmt.Println("Fibonacci Polynomial:", fibPolynomial)

	digits := "6287"
	combinations := algorithms.LetterCombinations(digits)
	fmt.Println("Letters Combinations:", combinations)

	growth := algorithms.ExponentialGrowth(10, 2, 4)
	fmt.Println("Exponential Growth:", growth)
}
