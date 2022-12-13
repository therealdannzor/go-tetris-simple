package main

import (
	"bufio"
	"os"
	"strings"
)

func readFromStdin() [][]string {
	var result [][]string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		token := scanner.Text()
		r := strings.Split(token, ",")
		result = append(result, r)
	}

	return result
}
