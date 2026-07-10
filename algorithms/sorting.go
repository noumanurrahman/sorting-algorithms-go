package algorithms

func BubbleSort(array []int) {
	swapping := true
	end := len(array) - 1
	for swapping {
		swapping = false
		for i := range end {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				swapping = true
			}
		}
	}
}

func InsertionSort(array []int) {
	i := 1
	j := 1
	for i < len(array) {
		if j > 0 && array[j] < array[j-1] {
			array[j], array[j-1] = array[j-1], array[j]
			j--
		} else {
			i++
			j = i
		}
	}
}

func SelectionSort(array []int) {
	for i := range array {
		smallest := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[smallest] {
				smallest = j
			}
		}
		array[i], array[smallest] = array[smallest], array[i]
	}
}
