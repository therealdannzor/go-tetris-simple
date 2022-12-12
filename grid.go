package main

import "errors"

// height is 0..100
type height int

const HEIGHT_LIMIT = 100

type Grid struct {
	fields           [HEIGHT_LIMIT]Row
	positionHeight   map[x_position]height
	currentMaxHeight height
}

type Shape struct {
	shortcode string // the uppercase letter representing the shape
	rows      int    // height taken up by the shape (e.g., I = 1, L = 4)
}

func (g *Grid) checkRowsToClear(index int) error {
	if g.fields[index].counter != 10 {
		return errors.New("row is not complete yet")
	}

	return nil
}

// Accepts a Shape and a position (0..9)
func (g *Grid) newShape(shape Shape, pos x_position) error {
	errBound := errors.New("invalid position, outside bound")
	errExisting := errors.New("existing shape overlap, cannot insert")

	switch shape.shortcode {
	case "Q":
		// Legend: S = existing, Q = new, pos = x, height = y
		// .  .  .  .  .
		// .  .  .  .  .
		// Q  Q  .  .  .
		// Q  Q  S  S  .
		// .  S  S  .  .

		if pos > 8 || pos < 0 {
			return errBound
		}
		h1 := g.positionHeight[pos]
		h2 := g.positionHeight[pos+1]
		h_max := maxHeight(h1, h2)
		g.fields[h_max].mark(pos)
		g.fields[h_max+1].mark(pos)
		g.fields[h_max].mark(pos + 1)
		g.fields[h_max+1].mark(pos + 1)

		// Q: height 2
		g.positionHeight[pos] = h_max + 2
		g.positionHeight[pos+1] = h_max + 2

		// potentially update the current max height of the whole grid
		g.currentMaxHeight = maxHeight(height(int(h_max)+2), g.currentMaxHeight)
		return nil

		//TODO: check if bottom row disappears
	case "Z":
		// Legend: I = existing, Z = new, pos = x, height = y
		// .  .  .  .  .
		// .  .  .  .  .
		// Z  Z  .  .  .
		// .  Z  Z  .  .
		// I  I  I  I  .

		if pos >= 8 || pos < 0 {
			return errBound
		}
		h1 := g.positionHeight[pos]
		h2 := g.positionHeight[pos+1]
		h3 := g.positionHeight[pos+2]
		h_max := maxHeight(h1, h2)
		h_max = maxHeight(h_max, h3)

		// Z: height 2
		g.positionHeight[pos] = h_max + 1
		g.positionHeight[pos+1] = h_max + 1
		g.positionHeight[pos+2] = h_max

		// nothing below us, just drop it!
		if h_max == 0 {
			g.fields[1].mark(pos)
			g.fields[1].mark(pos + 1)
			g.fields[0].mark(pos + 1)
			g.fields[0].mark(pos + 2)
			return nil
		}

		// the leftmost tile of the Z-shape will 'attach'
		if h_max == h1 {
			if g.fields[h_max].isFree(pos) && g.fields[h_max].isFree(pos+1) && g.fields[h_max-1].isFree(pos+1) && g.fields[h_max-1].isFree(pos+2) {
				g.fields[h_max].mark(pos)
				g.fields[h_max].mark(pos + 1)
				g.fields[h_max-1].mark(pos - 1)
				g.fields[h_max-1].mark(pos + 2)
			} else {
				return errExisting
			}
			// the middle two tiles of the Z-shape will 'attach'
		} else if h_max == h2 {
			if g.fields[h_max+1].isFree(pos-1) && g.fields[h_max+1].isFree(pos) && g.fields[h_max].isFree(pos) && g.fields[h_max].isFree(pos+1) {
				g.fields[h_max+1].mark(pos - 1)
				g.fields[h_max+1].mark(pos)
				g.fields[h_max].mark(pos)
				g.fields[h_max].mark(pos + 1)
			} else {
				return errExisting
			}
			// the rightmost tile
		} else if h_max == h3 {
			if g.fields[h_max+1].isFree(pos) && g.fields[h_max+1].isFree(pos+1) && g.fields[h_max].isFree(pos+1) && g.fields[h_max].isFree(pos+2) {
				g.fields[h_max+1].mark(pos)
				g.fields[h_max+1].mark(pos + 1)
				g.fields[h_max].mark(pos + 1)
				g.fields[h_max].mark(pos + 2)
			} else {
				return errExisting
			}
		}
		// potentially update the current max height of the whole grid
		//
		g.currentMaxHeight = maxHeight(height(int(h_max)+2), g.currentMaxHeight)
		return nil

	case "S":

		// CASE 1
		// Legend: L = existing, S = new, pos = x, height = y
		// .  S  S  .  .
		// S  S  .  .  .
		// L  .  .  .  .
		// L  .  .  .  .
		// L  L  .  .  .

		// CASE 2
		// Legend: T = existing, S = new, pos = x, height = y
		// .  .  .  .  .
		// .  .  S  S  .
		// .  S  S  .  .
		// .  .  T  T  T
		// .  .  .  T  .

		// CASE 3
		// Legend: L = existing, S = new, pos = x, height = y
		// .  .  .  .  .
		// .  .  S  S  .
		// .  S  S  L  .
		// .  .  .  L  .
		// .  .  .  L  L

		if pos >= 8 || pos < 0 {
			return errBound
		}

		h1 := g.positionHeight[pos]
		h2 := g.positionHeight[pos+1]
		h3 := g.positionHeight[pos+2]
		h_max := maxHeight(h1, h2)
		h_max = maxHeight(h_max, h3)

		// nothing below us, just drop it!
		if h_max == 0 {
			g.fields[0].mark(pos)
			g.fields[0].mark(pos + 1)
			g.fields[1].mark(pos + 1)
			g.fields[1].mark(pos + 2)

			// S: height 2
			g.positionHeight[pos] = 1
			g.positionHeight[pos+1] = 2
			g.positionHeight[pos+2] = 2

			// potentially update the current max height of the whole grid
			g.currentMaxHeight = maxHeight(height(int(h_max)+2), g.currentMaxHeight)
			return nil
		}

		// CASE 1
		// the leftmost tile of the S-shape touches an existing shape
		if h_max == h1 {
			if g.fields[h_max].isFree(pos) && g.fields[h_max].isFree(pos+1) && g.fields[h_max-1].isFree(pos+1) && g.fields[h_max-1].isFree(pos+2) {
				g.fields[h_max].mark(pos)
				g.fields[h_max].mark(pos + 1)
				g.fields[h_max-1].mark(pos - 1)
				g.fields[h_max-1].mark(pos + 2)

				// S: height 2
				g.positionHeight[pos] = h_max + 1
				g.positionHeight[pos+1] = h_max + 2
				g.positionHeight[pos+2] = h_max + 2
			} else {
				return errExisting
			}
			// CASE 2
			// the middle two tiles of the S-shape touches an existing shape
		} else if h_max == h2 {
			if g.fields[h_max+1].isFree(pos-1) && g.fields[h_max+1].isFree(pos) && g.fields[h_max].isFree(pos) && g.fields[h_max].isFree(pos+1) {
				g.fields[h_max+1].mark(pos - 1)
				g.fields[h_max+1].mark(pos)
				g.fields[h_max].mark(pos)
				g.fields[h_max].mark(pos + 1)

				// S: height 2
				g.positionHeight[pos+1] = h_max + 2
				g.positionHeight[pos] = h_max + 1
				g.positionHeight[pos+2] = h_max + 2

			} else {
				return errExisting
			}
			// CASE 3
			// the rightmost tile of the S-shape touches an existing shape
		} else if h_max == h3 {
			if g.fields[h_max-1].isFree(pos-2) && g.fields[h_max-1].isFree(pos-1) && g.fields[h_max].isFree(pos-1) && g.fields[h_max].isFree(pos) {
				g.fields[h_max-1].mark(pos - 2)
				g.fields[h_max-1].mark(pos - 1)
				g.fields[h_max].mark(pos - 1)
				g.fields[h_max].mark(pos)

				// S: height 2
				g.positionHeight[pos+2] = h_max + 1
				g.positionHeight[pos+1] = h_max + 1
				g.positionHeight[pos] = h_max

			} else {
				return errExisting
			}
		}

		// potentially update the current max height of the whole grid
		g.currentMaxHeight = maxHeight(height(int(h_max)+2), g.currentMaxHeight)
		return nil

	case "T":
		if pos >= 8 || pos < 0 {
			return errBound
		}

		h1 := g.positionHeight[pos]
		h2 := g.positionHeight[pos+1]
		h3 := g.positionHeight[pos+2]
		h_max := maxHeight(h1, h2)
		h_max = maxHeight(h_max, h3)

		// the leftmost tile of the S-shape will 'attach'
		if h_max == h1 {
			if g.fields[h_max].isFree(pos) && g.fields[h_max].isFree(pos+1) && g.fields[h_max-1].isFree(pos-1) && g.fields[h_max].isFree(pos+2) {
				g.fields[h_max].mark(pos)
				g.fields[h_max].mark(pos + 1)
				g.fields[h_max-1].mark(pos - 1)
				g.fields[h_max].mark(pos + 2)

				// T: height 2
				newMax := h_max + 1
				g.positionHeight[pos] = newMax
				g.positionHeight[pos+1] = newMax
				g.positionHeight[pos+2] = newMax

				// potentially update the current max height of the whole grid
				g.currentMaxHeight = maxHeight(newMax, g.currentMaxHeight)
			} else {
				return errExisting
			}
			// the middle two tiles of the S-shape will 'attach'
		} else if h_max == h2 {
			if g.fields[h_max].isFree(pos+1) && g.fields[h_max+1].isFree(pos+1) && g.fields[h_max+1].isFree(pos) && g.fields[h_max+1].isFree(pos+2) {
				g.fields[h_max].mark(pos + 1)
				g.fields[h_max+1].mark(pos + 1)
				g.fields[h_max+1].mark(pos)
				g.fields[h_max+1].mark(pos + 2)

				// T: height 2
				newMax := h_max + 2
				g.positionHeight[pos+1] = newMax
				g.positionHeight[pos] = newMax
				g.positionHeight[pos+2] = newMax

				// potentially update the current max height of the whole grid
				g.currentMaxHeight = maxHeight(newMax, g.currentMaxHeight)
			} else {
				return errExisting
			}
		} else if h_max == h3 {
			if g.fields[h_max].isFree(pos) && g.fields[h_max].isFree(pos+1) && g.fields[h_max-1].isFree(pos+1) && g.fields[h_max].isFree(pos+2) {
				g.fields[h_max].mark(pos)
				g.fields[h_max].mark(pos + 1)
				g.fields[h_max-1].mark(pos + 1)
				g.fields[h_max].mark(pos + 2)

				// T: height 2
				newMax := h_max + 1
				g.positionHeight[pos+1] = newMax
				g.positionHeight[pos] = newMax
				g.positionHeight[pos+2] = newMax

				// potentially update the current max height of the whole grid
				g.currentMaxHeight = maxHeight(newMax, g.currentMaxHeight)
			} else {
				return errExisting
			}
		}

		// potentially update the current max height of the whole grid
		g.currentMaxHeight = maxHeight(height(int(h_max)+2), g.currentMaxHeight)
		return nil

	case "I":
		if pos >= 7 || pos < 0 {
			return errBound
		}

		h1 := g.positionHeight[pos]
		h2 := g.positionHeight[pos+1]
		h3 := g.positionHeight[pos+2]
		h4 := g.positionHeight[pos+3]
		h_max := maxHeight(h1, h2)
		h_max = maxHeight(h_max, h3)
		h_max = maxHeight(h_max, h4)
		g.fields[h_max].mark(pos)
		g.fields[h_max].mark(pos + 1)
		g.fields[h_max].mark(pos + 2)
		g.fields[h_max].mark(pos + 3)

		// I: height 1
		newMax := h_max + 1
		g.positionHeight[pos] = newMax
		g.positionHeight[pos+1] = newMax
		g.positionHeight[pos+2] = newMax
		g.positionHeight[pos+3] = newMax
		// potentially update the current max height of the whole grid
		g.currentMaxHeight = maxHeight(newMax, g.currentMaxHeight)
		return nil

	case "L":
		if pos >= 9 || pos < 0 {
			return errBound
		}

		h1 := g.positionHeight[pos]
		h2 := g.positionHeight[pos+1]
		h_max := maxHeight(h1, h2)
		g.fields[h_max+2].mark(pos)
		g.fields[h_max+1].mark(pos)
		g.fields[h_max].mark(pos)
		g.fields[h_max].mark(pos + 1)

		// L: height 3
		newMax := h_max + 3
		g.positionHeight[pos] = newMax
		g.positionHeight[pos+1] += 1

		// potentially update the current max height of the whole grid
		g.currentMaxHeight = maxHeight(newMax, g.currentMaxHeight)
		return nil

	case "J":
		if pos >= 9 || pos < 0 {
			return errBound
		}

		h1 := g.positionHeight[pos]
		h2 := g.positionHeight[pos+1]
		h_max := maxHeight(h1, h2)
		g.fields[h_max].mark(pos)
		g.fields[h_max].mark(pos + 1)
		g.fields[h_max+1].mark(pos + 1)
		g.fields[h_max+2].mark(pos + 1)

		// J: height 3
		newMax := h_max + 3
		g.positionHeight[pos] += 1
		g.positionHeight[pos+1] = newMax

		// potentially update the current max height of the whole grid
		g.currentMaxHeight = maxHeight(newMax, g.currentMaxHeight)
		return nil

	}

	return nil
}

func maxHeight(h1 height, h2 height) height {
	if h2 >= h1 {
		return h2
	}
	return h1
}
