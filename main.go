package main

import (
	"algorithm-problems/two_sum"
	"fmt"
)

func main() {
	numbers := []int{2, 3, 5, 7, 9, 11}
	target := 9

	val1, val2 := two_sum.Two_Sum_HashTable(target, numbers)
	fmt.Println(val1, val2)
	fmt.Println()
}
