// Copyright 2026 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hilbert

// Sierpinski represents a 2D Sierpinski space of order N for mapping to and from.
// Implements SpaceFilling interface.
type Sierpinski struct {
	N int
}

// NewSierpinski returns a Sierpinski space which maps integers to and from the curve.
// n must be a power of two.
func NewSierpinski(n int) (*Sierpinski, error) {
	if n <= 0 {
		return nil, ErrNotPositive
	}

	// Test if power of two
	if (n & (n - 1)) != 0 {
		return nil, ErrNotPowerOfTwo
	}

	return &Sierpinski{
		N: n,
	}, nil
}

// GetDimensions returns the width and height of the 2D space.
func (s *Sierpinski) GetDimensions() (int, int) {
	return s.N, s.N
}

// GetGridType returns the geometry of the grid.
func (s *Sierpinski) GetGridType() GridType {
	return GridTriangular
}

// GetCount returns the total number of points on the curve.
func (s *Sierpinski) GetCount() int {
	return s.N * s.N
}

// Map transforms a one dimension value, t, in the range [0, n^2-1] to coordinates on the Sierpinski
// curve in a triangular grid, where y is the row index [0, n-1] and x is the triangle index
// in that row [0, 2y].
func (s *Sierpinski) Map(t int) (x, y int, err error) {
	if t < 0 || t >= s.N*s.N {
		return -1, -1, ErrOutOfRange
	}

	if s.N == 1 {
		return 0, 0, nil
	}

	// Calculate row y. Total triangles up to row y-1 is y^2.
	// So t = y^2 + x  =>  y = floor(sqrt(t))
	y = 0
	for (y+1)*(y+1) <= t {
		y++
	}

	x = t - y*y

	// Snake effect: alternate direction for even/odd rows to improve continuity
	if y%2 != 0 {
		x = (2 * y) - x
	}

	return x, y, nil
}

// MapInverse transform coordinates on Sierpinski curve from (x,y) to t.
func (s *Sierpinski) MapInverse(x, y int) (t int, err error) {
	if y < 0 || y >= s.N || x < 0 || x > 2*y {
		return -1, ErrOutOfRange
	}

	if y%2 != 0 {
		x = (2 * y) - x
	}

	return y*y + x, nil
}

