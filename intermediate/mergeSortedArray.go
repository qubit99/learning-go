package main

import (
	"fmt"

	util "example.com/intermediate/util/pkg"
)

func MergeSortedArrays(arr1, arr2 []int) []int {
	i, j := 0, 0
	merged := make([]int, 0, len(arr1)+len(arr2))

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			merged = append(merged, arr1[i])
			i++
		} else {
			merged = append(merged, arr2[j])
			j++
		}
	}

	// Append any remaining elements from arr1 or arr2
	for i < len(arr1) {
		merged = append(merged, arr1[i])
		i++
	}
	for j < len(arr2) {
		merged = append(merged, arr2[j])
		j++
	}

	return merged
}

func main() {
	arr1, _ := util.ReadIntegerArrayFromFile("array1")
	arr2, _ := util.ReadIntegerArrayFromFile("array2")
	merged := MergeSortedArrays(arr1, arr2)
	fmt.Println("Merged array:", merged)
}
