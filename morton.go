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

// Morton represents a 2D Morton (Z-order) curve of order N for mapping to and from.
// Implements SpaceFilling interface.
type Morton struct {
	N int
}

// NewMorton returns a Morton space which maps integers to and from the curve.
// n must be a power of two.
func NewMorton(n int) (*Morton, error) {
	if n <= 0 {
		return nil, ErrNotPositive
	}

	// Test if power of two
	if (n & (n - 1)) != 0 {
		return nil, ErrNotPowerOfTwo
	}

	return &Morton{
		N: n,
	}, nil
}

// GetDimensions returns the width and height of the 2D space.
func (m *Morton) GetDimensions() (int, int) {
	return m.N, m.N
}

// GetGridType returns the geometry of the grid.
func (m *Morton) GetGridType() GridType {
	return GridSquare
}

// GetCount returns the total number of points on the curve.
func (m *Morton) GetCount() int {
	return m.N * m.N
}

// Map transforms a one dimension value, t, in the range [0, n^2-1] to coordinates on the Morton
// curve in the two-dimension space, where x and y are within [0,n-1].
func (m *Morton) Map(t int) (x, y int, err error) {
	if t < 0 || t >= m.N*m.N {
		return -1, -1, ErrOutOfRange
	}

	for i := 0; (1 << i) < m.N; i++ {
		x |= ((t >> (2 * i + 1)) & 1) << i
		y |= ((t >> (2 * i)) & 1) << i
	}

	return x, y, nil
}

// MapInverse transform coordinates on Morton curve from (x,y) to t.
func (m *Morton) MapInverse(x, y int) (t int, err error) {
	if x < 0 || x >= m.N || y < 0 || y >= m.N {
		return -1, ErrOutOfRange
	}

	for i := 0; (1 << i) < m.N; i++ {
		t |= ((x >> i) & 1) << (2 * i + 1)
		t |= ((y >> i) & 1) << (2 * i)
	}

	return t, nil
}
