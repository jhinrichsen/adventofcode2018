package adventofcode2018

import (
	"image"
	"iter"
)

// Grid defines board dimensions.
type Grid struct {
	W, H int
}

// C6Indices yields (index, neighbors) for 6-connectivity (hexagonal grid).
// Uses axial coordinates: E(+1), W(-1), NE(+W), SW(-W), NW(+W-1), SE(-W+1)
// Row 0 is bottom, row H-1 is top.
func (g Grid) C6Indices() iter.Seq2[int, iter.Seq[int]] {
	w, h := g.W, g.H

	return func(yield func(int, iter.Seq[int]) bool) {
		// Bottom-left corner (idx=0): E, NE (2 neighbors)
		if !yield(0, func(y func(int) bool) {
			_ = y(1) && y(w)
		}) {
			return
		}

		// Bottom-right corner: W, NE, NW (3 neighbors)
		br0 := w - 1
		if !yield(br0, func(y func(int) bool) {
			_ = y(br0-1) && y(br0+w) && y(br0+w-1)
		}) {
			return
		}

		// Top-left corner: E, SW, SE (3 neighbors)
		tl := (h - 1) * w
		if !yield(tl, func(y func(int) bool) {
			_ = y(tl+1) && y(tl-w) && y(tl-w+1)
		}) {
			return
		}

		// Top-right corner: W, SW (2 neighbors)
		tr := h*w - 1
		if !yield(tr, func(y func(int) bool) {
			_ = y(tr-1) && y(tr-w)
		}) {
			return
		}

		// Bottom edge (row 0): E, W, NE, NW (4 neighbors, no SW/SE)
		for x := 1; x < w-1; x++ {
			idx := x
			if !yield(idx, func(y func(int) bool) {
				_ = y(idx+1) && y(idx-1) && y(idx+w) && y(idx+w-1)
			}) {
				return
			}
		}

		// Top edge (row H-1): E, W, SW, SE (4 neighbors, no NE/NW)
		for x := 1; x < w-1; x++ {
			idx := (h-1)*w + x
			if !yield(idx, func(y func(int) bool) {
				_ = y(idx+1) && y(idx-1) && y(idx-w) && y(idx-w+1)
			}) {
				return
			}
		}

		// Left edge (col 0): E, NE, SW, SE (4 neighbors, no W/NW)
		for row := 1; row < h-1; row++ {
			idx := row * w
			if !yield(idx, func(y func(int) bool) {
				_ = y(idx+1) && y(idx+w) && y(idx-w) && y(idx-w+1)
			}) {
				return
			}
		}

		// Right edge (col W-1): W, NE, NW, SW (4 neighbors, no E/SE)
		for row := 1; row < h-1; row++ {
			idx := row*w + w - 1
			if !yield(idx, func(y func(int) bool) {
				_ = y(idx-1) && y(idx+w) && y(idx+w-1) && y(idx-w)
			}) {
				return
			}
		}

		// Interior (6 neighbors each)
		for row := 1; row < h-1; row++ {
			for col := 1; col < w-1; col++ {
				idx := row*w + col
				if !yield(idx, func(y func(int) bool) {
					_ = y(idx+1) && y(idx-1) && // E, W
						y(idx+w) && y(idx-w) && // NE, SW
						y(idx+w-1) && y(idx-w+1) // NW, SE
				}) {
					return
				}
			}
		}
	}
}

// C4Indices yields (index, neighbors) for 4-connectivity.
func (g Grid) C4Indices() iter.Seq2[int, iter.Seq[int]] {
	w, h := g.W, g.H

	return func(yield func(int, iter.Seq[int]) bool) {
		// Top-left corner (2 neighbors)
		if !yield(0, func(y func(int) bool) {
			_ = y(1) && y(w)
		}) {
			return
		}

		// Top-right corner (2 neighbors)
		tr := w - 1
		if !yield(tr, func(y func(int) bool) {
			_ = y(tr-1) && y(tr+w)
		}) {
			return
		}

		// Bottom-left corner (2 neighbors)
		bl := (h - 1) * w
		if !yield(bl, func(y func(int) bool) {
			_ = y(bl-w) && y(bl+1)
		}) {
			return
		}

		// Bottom-right corner (2 neighbors)
		br := h*w - 1
		if !yield(br, func(y func(int) bool) {
			_ = y(br-w) && y(br-1)
		}) {
			return
		}

		// Top edge (3 neighbors each)
		for idx := 1; idx < w-1; idx++ {
			if !yield(idx, func(y func(int) bool) {
				_ = y(idx-1) && y(idx+1) && y(idx+w)
			}) {
				return
			}
		}

		// Bottom edge (3 neighbors each)
		for x := 1; x < w-1; x++ {
			idx := (h-1)*w + x
			if !yield(idx, func(y func(int) bool) {
				_ = y(idx-w) && y(idx-1) && y(idx+1)
			}) {
				return
			}
		}

		// Left edge (3 neighbors each)
		for row := 1; row < h-1; row++ {
			idx := row * w
			if !yield(idx, func(y func(int) bool) {
				_ = y(idx-w) && y(idx+1) && y(idx+w)
			}) {
				return
			}
		}

		// Right edge (3 neighbors each)
		for row := 1; row < h-1; row++ {
			idx := row*w + w - 1
			if !yield(idx, func(y func(int) bool) {
				_ = y(idx-w) && y(idx-1) && y(idx+w)
			}) {
				return
			}
		}

		// Interior (4 neighbors each)
		for row := 1; row < h-1; row++ {
			for col := 1; col < w-1; col++ {
				idx := row*w + col
				if !yield(idx, func(y func(int) bool) {
					_ = y(idx-w) && y(idx-1) && y(idx+1) && y(idx+w)
				}) {
					return
				}
			}
		}
	}
}

// C8Indices yields (index, neighbors) for 8-connectivity.
func (g Grid) C8Indices() iter.Seq2[int, iter.Seq[int]] {
	w, h := g.W, g.H

	return func(yield func(int, iter.Seq[int]) bool) {
		// Top-left corner (3 neighbors)
		if !yield(0, func(y func(int) bool) {
			_ = y(1) && y(w) && y(w+1)
		}) {
			return
		}

		// Top-right corner (3 neighbors)
		tr := w - 1
		if !yield(tr, func(y func(int) bool) {
			_ = y(tr-1) && y(tr+w-1) && y(tr+w)
		}) {
			return
		}

		// Bottom-left corner (3 neighbors)
		bl := (h - 1) * w
		if !yield(bl, func(y func(int) bool) {
			_ = y(bl-w) && y(bl-w+1) && y(bl+1)
		}) {
			return
		}

		// Bottom-right corner (3 neighbors)
		br := h*w - 1
		if !yield(br, func(y func(int) bool) {
			_ = y(br-w-1) && y(br-w) && y(br-1)
		}) {
			return
		}

		// Top edge (5 neighbors each)
		for idx := 1; idx < w-1; idx++ {
			if !yield(idx, func(y func(int) bool) {
				_ = y(idx-1) && y(idx+1) && y(idx+w-1) && y(idx+w) && y(idx+w+1)
			}) {
				return
			}
		}

		// Bottom edge (5 neighbors each)
		for x := 1; x < w-1; x++ {
			idx := (h-1)*w + x
			if !yield(idx, func(y func(int) bool) {
				_ = y(idx-w-1) && y(idx-w) && y(idx-w+1) && y(idx-1) && y(idx+1)
			}) {
				return
			}
		}

		// Left edge (5 neighbors each)
		for row := 1; row < h-1; row++ {
			idx := row * w
			if !yield(idx, func(y func(int) bool) {
				_ = y(idx-w) && y(idx-w+1) && y(idx+1) && y(idx+w) && y(idx+w+1)
			}) {
				return
			}
		}

		// Right edge (5 neighbors each)
		for row := 1; row < h-1; row++ {
			idx := row*w + w - 1
			if !yield(idx, func(y func(int) bool) {
				_ = y(idx-w-1) && y(idx-w) && y(idx-1) && y(idx+w-1) && y(idx+w)
			}) {
				return
			}
		}

		// Interior (8 neighbors each)
		for row := 1; row < h-1; row++ {
			for col := 1; col < w-1; col++ {
				idx := row*w + col
				if !yield(idx, func(y func(int) bool) {
					_ = y(idx-w-1) && y(idx-w) && y(idx-w+1) &&
						y(idx-1) && y(idx+1) &&
						y(idx+w-1) && y(idx+w) && y(idx+w+1)
				}) {
					return
				}
			}
		}
	}
}

// C4Points yields (pos, neighbors) for 4-connectivity.
func (g Grid) C4Points() iter.Seq2[image.Point, iter.Seq[image.Point]] {
	w, h := g.W, g.H

	return func(yield func(image.Point, iter.Seq[image.Point]) bool) {
		// Top-left corner (2 neighbors)
		if !yield(image.Point{0, 0}, func(y func(image.Point) bool) {
			_ = y(image.Point{1, 0}) && y(image.Point{0, 1})
		}) {
			return
		}

		// Top-right corner (2 neighbors)
		if !yield(image.Point{w - 1, 0}, func(y func(image.Point) bool) {
			_ = y(image.Point{w - 2, 0}) && y(image.Point{w - 1, 1})
		}) {
			return
		}

		// Bottom-left corner (2 neighbors)
		if !yield(image.Point{0, h - 1}, func(y func(image.Point) bool) {
			_ = y(image.Point{0, h - 2}) && y(image.Point{1, h - 1})
		}) {
			return
		}

		// Bottom-right corner (2 neighbors)
		if !yield(image.Point{w - 1, h - 1}, func(y func(image.Point) bool) {
			_ = y(image.Point{w - 1, h - 2}) && y(image.Point{w - 2, h - 1})
		}) {
			return
		}

		// Top edge (3 neighbors each)
		for x := 1; x < w-1; x++ {
			if !yield(image.Point{x, 0}, func(y func(image.Point) bool) {
				_ = y(image.Point{x - 1, 0}) && y(image.Point{x + 1, 0}) && y(image.Point{x, 1})
			}) {
				return
			}
		}

		// Bottom edge (3 neighbors each)
		for x := 1; x < w-1; x++ {
			if !yield(image.Point{x, h - 1}, func(y func(image.Point) bool) {
				_ = y(image.Point{x, h - 2}) && y(image.Point{x - 1, h - 1}) && y(image.Point{x + 1, h - 1})
			}) {
				return
			}
		}

		// Left edge (3 neighbors each)
		for row := 1; row < h-1; row++ {
			if !yield(image.Point{0, row}, func(y func(image.Point) bool) {
				_ = y(image.Point{0, row - 1}) && y(image.Point{1, row}) && y(image.Point{0, row + 1})
			}) {
				return
			}
		}

		// Right edge (3 neighbors each)
		for row := 1; row < h-1; row++ {
			if !yield(image.Point{w - 1, row}, func(y func(image.Point) bool) {
				_ = y(image.Point{w - 1, row - 1}) && y(image.Point{w - 2, row}) && y(image.Point{w - 1, row + 1})
			}) {
				return
			}
		}

		// Interior (4 neighbors each)
		for row := 1; row < h-1; row++ {
			for col := 1; col < w-1; col++ {
				if !yield(image.Point{col, row}, func(y func(image.Point) bool) {
					_ = y(image.Point{col, row - 1}) && y(image.Point{col - 1, row}) && y(image.Point{col + 1, row}) && y(image.Point{col, row + 1})
				}) {
					return
				}
			}
		}
	}
}

// C8Points yields (pos, neighbors) for 8-connectivity.
func (g Grid) C8Points() iter.Seq2[image.Point, iter.Seq[image.Point]] {
	w, h := g.W, g.H

	return func(yield func(image.Point, iter.Seq[image.Point]) bool) {
		// Top-left corner (3 neighbors)
		if !yield(image.Point{0, 0}, func(y func(image.Point) bool) {
			_ = y(image.Point{1, 0}) && y(image.Point{0, 1}) && y(image.Point{1, 1})
		}) {
			return
		}

		// Top-right corner (3 neighbors)
		if !yield(image.Point{w - 1, 0}, func(y func(image.Point) bool) {
			_ = y(image.Point{w - 2, 0}) && y(image.Point{w - 2, 1}) && y(image.Point{w - 1, 1})
		}) {
			return
		}

		// Bottom-left corner (3 neighbors)
		if !yield(image.Point{0, h - 1}, func(y func(image.Point) bool) {
			_ = y(image.Point{0, h - 2}) && y(image.Point{1, h - 2}) && y(image.Point{1, h - 1})
		}) {
			return
		}

		// Bottom-right corner (3 neighbors)
		if !yield(image.Point{w - 1, h - 1}, func(y func(image.Point) bool) {
			_ = y(image.Point{w - 2, h - 2}) && y(image.Point{w - 1, h - 2}) && y(image.Point{w - 2, h - 1})
		}) {
			return
		}

		// Top edge (5 neighbors each)
		for x := 1; x < w-1; x++ {
			if !yield(image.Point{x, 0}, func(y func(image.Point) bool) {
				_ = y(image.Point{x - 1, 0}) && y(image.Point{x + 1, 0}) &&
					y(image.Point{x - 1, 1}) && y(image.Point{x, 1}) && y(image.Point{x + 1, 1})
			}) {
				return
			}
		}

		// Bottom edge (5 neighbors each)
		for x := 1; x < w-1; x++ {
			if !yield(image.Point{x, h - 1}, func(y func(image.Point) bool) {
				_ = y(image.Point{x - 1, h - 2}) && y(image.Point{x, h - 2}) && y(image.Point{x + 1, h - 2}) &&
					y(image.Point{x - 1, h - 1}) && y(image.Point{x + 1, h - 1})
			}) {
				return
			}
		}

		// Left edge (5 neighbors each)
		for row := 1; row < h-1; row++ {
			if !yield(image.Point{0, row}, func(y func(image.Point) bool) {
				_ = y(image.Point{0, row - 1}) && y(image.Point{1, row - 1}) &&
					y(image.Point{1, row}) &&
					y(image.Point{0, row + 1}) && y(image.Point{1, row + 1})
			}) {
				return
			}
		}

		// Right edge (5 neighbors each)
		for row := 1; row < h-1; row++ {
			if !yield(image.Point{w - 1, row}, func(y func(image.Point) bool) {
				_ = y(image.Point{w - 2, row - 1}) && y(image.Point{w - 1, row - 1}) &&
					y(image.Point{w - 2, row}) &&
					y(image.Point{w - 2, row + 1}) && y(image.Point{w - 1, row + 1})
			}) {
				return
			}
		}

		// Interior (8 neighbors each)
		for row := 1; row < h-1; row++ {
			for col := 1; col < w-1; col++ {
				if !yield(image.Point{col, row}, func(y func(image.Point) bool) {
					_ = y(image.Point{col - 1, row - 1}) && y(image.Point{col, row - 1}) && y(image.Point{col + 1, row - 1}) &&
						y(image.Point{col - 1, row}) && y(image.Point{col + 1, row}) &&
						y(image.Point{col - 1, row + 1}) && y(image.Point{col, row + 1}) && y(image.Point{col + 1, row + 1})
				}) {
					return
				}
			}
		}
	}
}
