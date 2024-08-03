package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(canBeEqual([]int{1, 2, 3, 4}, []int{2, 4, 1, 3}))
}
func canBeEqual(target []int, arr []int) bool {
	sort.Ints(target)
	sort.Ints(arr)
	for i := range target {
		if target[i] != arr[i] {
			return false

		}
	}
	return true
}
