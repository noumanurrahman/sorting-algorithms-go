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
