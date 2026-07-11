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

func MergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	mid := len(array) / 2

	sortedFirst := MergeSort(array[:mid])
	sortedSecond := MergeSort(array[mid:])

	return merge(sortedFirst, sortedSecond)
}

func merge(first []int, second []int) []int {
	i := 0
	j := 0

	var final []int

	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			final = append(final, first[i])
			i++
		} else {
			final = append(final, second[j])
			j++
		}
	}

	for i < len(first) {
		final = append(final, first[i])
		i++
	}
	for j < len(second) {
		final = append(final, second[j])
		j++
	}

	return final
}
