package main

import "fmt"

func main() {

	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if num <= 1 {
		fmt.Printf("%d is not a prime number.\n", num)
		return
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			fmt.Printf("%d is not a prime number.\n", num)
			return
		}
	}
	fmt.Printf("%d is a prime number.\n", num)
}
