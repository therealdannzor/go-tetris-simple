package main

import "errors"

// height is 0..100
type height int

const MAX_HEIGHT = 100

type Grid struct {
	fields         [MAX_HEIGHT]Row
	positionHeight map[position]height
}

type Shape struct {
	shortcode string // the uppercase letter representing the shape
	rows      int    // height taken up by the shape (e.g., I = 1, L = 4)
}

// Accepts a Shape and a position (0..9)
func (g *Grid) newShape(shape Shape, pos position) error {
	errBound := errors.New("invalid position, outside bound")
	errExisting := errors.New("existing shape overlap, cannot insert")

	switch shape.shortcode {
	case "Q":
		if pos > 8 {
			return errBound
		}
		h1 := g.positionHeight[pos]
		h2 := g.positionHeight[pos+1]   // Legend: S = existing, Q = new, pos = x, height = y
		h_max := maxHeight(h1, h2)      // .  .  .  .  .
		g.fields[h_max].mark(pos)       // .  .  .  .  .
		g.fields[h_max+1].mark(pos)     // Q  Q  .  .  .
		g.fields[h_max].mark(pos + 1)   // Q  Q  S  S  .
		g.fields[h_max+1].mark(pos + 1) // .  S  S  .  .

		// the Q shape has height 2
		g.positionHeight[pos] = h_max + 2
		g.positionHeight[pos+1] = h_max + 2

		//TODO: check if bottom row disappears
	case "Z":
		if pos >= 8 {
			return errBound
		}
		h1 := g.positionHeight[pos]
		h2 := g.positionHeight[pos+1] // Legend: T = existing, Z = new
		h3 := g.positionHeight[pos+2]
		h_max := maxHeight(h1, h2)
		h_max = maxHeight(h_max, h3)

		// nothing below us, just drop it!
		if h_max == 0 {
			g.fields[1].mark(pos)
			g.fields[1].mark(pos + 1)
			g.fields[0].mark(pos + 1)
			g.fields[0].mark(pos + 2)
		}

		// the leftmost tile of the Z-shape will 'attach'
		if h_max == h1 {
			if !g.fields[h_max].isFree(pos) && !g.fields[h_max].isFree(pos+1) && !g.fields[h_max-1].isFree(pos+1) && !g.fields[h_max-1].isFree(pos+2) {
				g.fields[h_max].mark(pos)
				g.fields[h_max].mark(pos + 1)
				g.fields[h_max-1].mark(pos - 1)
				g.fields[h_max-1].mark(pos + 2)
			} else {
				return errExisting
			}
			// the middle two tiles of the Z-shape will 'attach'
		} else if h_max == h2 {
			if !g.fields[h_max+1].isFree(pos-1) && !g.fields[h_max+1].isFree(pos) && !g.fields[h_max].isFree(pos) && !g.fields[h_max].isFree(pos+1) {
				g.fields[h_max+1].mark(pos - 1)
				g.fields[h_max+1].mark(pos)
				g.fields[h_max].mark(pos)
				g.fields[h_max].mark(pos + 1)
			} else {
				return errExisting
			}
		} else if h_max == h3 {
			if !g.fields[h_max+1].isFree(pos) && !g.fields[h_max+1].isFree(pos+1) && !g.fields[h_max].isFree(pos+1) && !g.fields[h_max].isFree(pos+2) {
				g.fields[h_max+1].mark(pos)
				g.fields[h_max+1].mark(pos + 1)
				g.fields[h_max].mark(pos + 1)
				g.fields[h_max].mark(pos + 2)
			} else {
				return errExisting
			}
		}

	case "S":
		// mark two rows
	case "T":
		// mark two rows
	case "I":
		// mark one row
	case "L":
		// mark four rows
	case "J":
		// mark four rows

	}

	return nil
}

func maxHeight(h1 height, h2 height) height {
	if h2 >= h1 {
		return h2
	}
	return h1
}
