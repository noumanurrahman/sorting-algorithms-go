package algorithms

import (
	"log"
)

func FibonacciExponential(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return FibonacciExponential(n-1) + FibonacciExponential(n-2)
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	current := 0
	parent := 1
	grandparent := 0

	for _ = range n - 1 {
		current = parent + grandparent
		grandparent = parent
		parent = current
	}

	return current
}

func LetterCombinations(digits string) []string {
	combinations := []string{""}
	for _, digit := range digits {
		letters, ok := digitsToLetters[string(digit)]
		if ok {
			var newResult []string
			for _, combo := range combinations {
				for _, letter := range letters {
					newResult = append(newResult, combo+string(letter))
				}
			}
			combinations = newResult
		} else {
			log.Fatal("letters not found in digits to letters")
		}
	}
	return combinations
}

var digitsToLetters = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func ExponentialGrowth(initial int, factor int, time int) []int {
	sequence := []int{initial}
	for i := 1; i <= time; i++ {
		sequence = append(sequence, sequence[i-1]*factor)
	}
	return sequence
}
