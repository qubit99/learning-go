package main

import "fmt"

func main() {
	a, b := 0, 1
	for i := 0; i < 20; i++ {
		fmt.Printf("%d ", a)
		a, b = b, a+b
	}
	fmt.Println()
}
