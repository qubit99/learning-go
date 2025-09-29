package main

import "fmt"

func modify(arr *[5]int) {
	arr[0] = 42
}

func main() {
	var arr1 = [3]int{1, 2, 3}
	c := arr1[0]
	c = c + 1
	var arr2 [5]int
	fmt.Println("arr2 before modification:", arr2)
	modify(&arr2)
	fmt.Println("arr2 after modification:", arr2)
}
