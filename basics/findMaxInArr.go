package main

import "fmt"

func main() {
	var arr [100]int
	fmt.Print("Enter numbers (space-separated): ")
	for i := 0; i < 100; i++ {
		_, err := fmt.Scan(&arr[i])
		if err != nil {
			break
		}
	}

	max := arr[0]
	for i := 1; i < 100; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	fmt.Printf("The maximum number is: %d\n", max)
}
