package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadShapes(t *testing.T) {
	testCases := [][]string{
		{"I0", "I4", "Q8"},
		{"T1", "Z3", "I4"},
		{"Q0", "I2", "I6", "I0", "I6", "I6", "Q2", "Q4"},
		{"L1", "L1", "L1"},
		{"Q0", "Q2", "Q4", "Q6", "Q8"},
		{"S0", "S2", "S4", "S6", "S7"},
		{"Q0"},
		{"Q0", "Q1"},
		{"I0", "I4", "Q8"},
		{"I0", "I4", "Q8", "I0", "I4"},
		{"L0", "J2", "L4", "J6", "Q8"},
		{"L0", "Z1", "Z3", "Z5", "Z7"},
		{"T0", "T3"},
		{"T0", "T3", "I6", "I6"},
		{"I0", "I6", "S4"},
		{"T1", "Z3", "I4"},
		{"L0", "J3", "L5", "J8", "T1"},
		{"L0", "J3", "L5", "J8", "T1", "T6"},
		{"L0", "J3", "L5", "J8", "T1", "T6", "J2", "L6", "T0", "T7"},
		{"L0", "J3", "L5", "J8", "T1", "T6", "J2", "L6", "T0", "T7", "Q4"},
		{"S0", "S2", "S4", "S6"},
		{"S0", "S2", "S4", "S5", "Q8", "Q8", "Q8", "Q8", "T1", "Q1", "I0","Q4"},
		{"L0", "J3", "L5", "J8", "T1", "T6", "S2", "Z5", "T0", "T7"},
		{"Q0","I2","I6","I0","I6","I6","Q2","Q4"},
	}
	expected := []int{1, 4, 3, 9, 0, 10, 2, 4, 1, 0, 2, 2, 2, 1, 1, 4, 3, 1, 2, 1, 8, 8, 0, 3}

	for i := range testCases {
		grid := NewGrid()
		have := grid.ReadShapes(testCases[i])
		want := expected[i]
		assert.Equal(t, want, have)
	}

}
