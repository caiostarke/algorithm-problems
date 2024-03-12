package two_sum

// Given an ordered array of integers find the sum of two numbers
// O(n)
// Return the values, not indexes
func Two_sum(target int, array []int) (int, int) {
	// two pointers
	left := 0
	right := len(array) - 1

	for left <= right {
		if array[left]+array[right] == target {
			return array[left], array[right]
		} else if array[left]+array[right] > target {
			right--
		} else {
			left++
		}
	}

	return -1, -1
}
