package main

import (
	"algorithm-problems/binary_tree"
	"fmt"
)

func main() {
	bt := &binary_tree.BynaryTree{}

	values := []int{4, 2, 7, 1, 3, 6, 9}

	for _, v := range values {
		bt.Insert(v)
	}

	fmt.Println("Maximum value in the binary tree:", bt.FindMax())
}

//		4
//	 2   	7
// 1   3  6   9
//
