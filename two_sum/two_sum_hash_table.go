package two_sum

import "fmt"

// target = 9
// [ 1, 2, 4, 6, 7, 11]
// hashTable = {
// 8: 0
// 7: 1
// 5: 2
// 3: 3
// }
func Two_Sum_HashTable(target int, numbers []int) (int, int) {
	hashtable := map[int]int{}

	for i := 0; i < len(numbers); i++ {
		index := target - numbers[i]

		if val, ok := hashtable[numbers[i]]; ok {
			return numbers[val], numbers[i]
		}

		fmt.Println(hashtable)

		hashtable[index] = i
	}

	return -1, -1
}
