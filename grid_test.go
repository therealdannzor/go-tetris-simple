package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadShapes(t *testing.T) {
	testCases := [][]string{{"I0", "I4", "Q8"}, {"T1", "Z3", "I4"}, {"Q0", "I2", "I6", "I0", "I6", "I6", "Q2", "Q4"},
	{"L1", "L1", "L1"}, {"Q0", "Q2", "Q4", "Q6", "Q8"}, {"S0", "S2", "S4", "S6", "S7"}}
	expected := []int{1, 4, 3, 9, 0, 10}

	for i := range testCases {
		grid := NewGrid()
		have := grid.ReadShapes(testCases[i])
		want := expected[i]
		assert.Equal(t, want, have)
	}

}
