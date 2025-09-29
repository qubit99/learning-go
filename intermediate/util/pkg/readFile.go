package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadIntegerArrayFromFile(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := bufio.NewReader(file)
	line, err := r.ReadString('\n') // if file has no trailing newline, err can be io.EOF
	if err != nil && len(line) == 0 {
		return nil, err // nothing read and got an error
	}

	fields := strings.Fields(line) // splits on any whitespace
	ints := make([]int, 0, len(fields))
	for _, tok := range fields {
		n, err := strconv.Atoi(tok)
		if err != nil {
			return nil, fmt.Errorf("bad int %q: %w", tok, err)
		}
		ints = append(ints, n)
	}
	return ints, nil
}
