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

// Snake represents a 2D Snake (Boustrophedon) scan of order N.
// Implements SpaceFilling2D interface.
type Snake struct {
	N int
}

// NewSnake returns a Snake space which maps integers to and from the curve.
func NewSnake(n int) (*Snake, error) {
	if n <= 0 {
		return nil, ErrNotPositive
	}
	return &Snake{N: n}, nil
}

// GetDimensions returns the width and height of the 2D space.
func (s *Snake) GetDimensions() (int, int) {
	return s.N, s.N
}

// GetGridType returns the geometry of the grid.
func (s *Snake) GetGridType() GridType {
	return GridSquare
}

// GetCount returns the total number of points on the curve.
func (s *Snake) GetCount() int {
	return s.N * s.N
}

// Map transforms a one dimension value, t, in the range [0, n^2-1] to coordinates on the Snake
// scan in the two-dimension space, where x and y are within [0,n-1].
func (s *Snake) Map(t int) (x, y int, err error) {
	if t < 0 || t >= s.N*s.N {
		return -1, -1, ErrOutOfRange
	}

	y = t / s.N
	x = t % s.N

	if y%2 != 0 {
		x = s.N - 1 - x
	}

	return x, y, nil
}

// MapInverse transform coordinates on Snake scan from (x,y) to t.
func (s *Snake) MapInverse(x, y int) (t int, err error) {
	if x < 0 || x >= s.N || y < 0 || y >= s.N {
		return -1, ErrOutOfRange
	}

	t = y * s.N
	if y%2 == 0 {
		t += x
	} else {
		t += (s.N - 1 - x)
	}

	return t, nil
}
