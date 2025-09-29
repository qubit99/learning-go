package main

import "fmt"

func main() {
	var word string
	fmt.Print("Enter a word: ")
	fmt.Scan(&word)

	n := len(word)
	for i := 0; i < n/2; i++ {
		if word[i] != word[n-1-i] {
			fmt.Printf("%s is not a palindrome.\n", word)
			return
		}
	}
	fmt.Printf("%s is a palindrome.\n", word)
}
