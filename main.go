package main

import "fmt"

func main() {
	lines := readFromStdin()

	for i := range lines {
		grid := NewGrid()
		result := grid.ReadShapes(lines[i])
		fmt.Println(result)
	}
}
