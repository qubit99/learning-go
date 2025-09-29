package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter space-separated ints: ")
	line, _ := reader.ReadString('\n') // read one line (includes '\n')
	fields := strings.Fields(line)     // split on any whitespace
	nums := make([]int, len(fields))   // slice sized to input

	for i, f := range fields {
		n, err := strconv.Atoi(f) // string -> int
		if err != nil {
			fmt.Printf("bad int %q: %v\n", f, err)
			return
		}
		nums[i] = n
	}

	sum := 0
	for _, n := range nums {
		sum += n
	}

	fmt.Printf("The sum is: %d\n", sum)
}
