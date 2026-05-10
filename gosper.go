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

import (
	"math"
)

// Gosper represents a 2D Gosper curve of order N.
// The Gosper curve is a space-filling curve on a hexagonal grid.
// Implements SpaceFilling2D interface.
type Gosper struct {
	Order int
	N     int // Number of hexagons = 7^Order
	Path  [][2]int
}

// NewGosper returns a Gosper space which maps integers to and from the curve.
// order is the recursive depth. Total hexagons will be 7^order.
func NewGosper(order int) (*Gosper, error) {
	if order < 0 {
		return nil, ErrNotPositive
	}

	n := int(math.Pow(7, float64(order)))
	g := &Gosper{
		Order: order,
		N:     n,
		Path:  make([][2]int, n),
	}

	g.generatePath()
	return g, nil
}

func (g *Gosper) generatePath() {
	// Directions in axial coordinates (q, r)
	dirs := [6][2]int{
		{1, 0}, {1, -1}, {0, -1}, {-1, 0}, {-1, 1}, {0, 1},
	}

	q, r := 0, 0
	angle := 0
	idx := 0

	var a, b func(n int)
	move := func() {
		if idx < g.N {
			g.Path[idx] = [2]int{q, r}
			idx++

			ang := ((angle % 6) + 6) % 6
			q += dirs[ang][0]
			r += dirs[ang][1]
		}
	}

	a = func(n int) {
		if n == 0 {
			move()
			return
		}
		a(n - 1)
		angle--
		b(n - 1)
		angle -= 2
		b(n - 1)
		angle++
		a(n - 1)
		angle += 2
		a(n - 1)
		a(n - 1)
		angle++
		b(n - 1)
		angle--
	}

	b = func(n int) {
		if n == 0 {
			move()
			return
		}
		angle++
		a(n - 1)
		angle--
		b(n - 1)
		b(n - 1)
		angle -= 2
		b(n - 1)
		angle--
		a(n - 1)
		angle += 2
		a(n - 1)
		angle++
		b(n - 1)
	}

	a(g.Order)
}

// GetDimensions returns the width and height of the bounding box for the hexagonal grid.
func (g *Gosper) GetDimensions() (int, int) {
	if len(g.Path) == 0 {
		return 0, 0
	}
	minQ, maxQ, minR, maxR := 0, 0, 0, 0
	for _, p := range g.Path {
		if p[0] < minQ { minQ = p[0] }
		if p[0] > maxQ { maxQ = p[0] }
		if p[1] < minR { minR = p[1] }
		if p[1] > maxR { maxR = p[1] }
	}
	return maxQ - minQ + 1, maxR - minR + 1
}

// GetGridType returns the geometry of the grid.
func (g *Gosper) GetGridType() GridType {
	return GridHexagonal
}

// GetCount returns the total number of points on the curve.
func (g *Gosper) GetCount() int {
	return g.N
}

// Map transforms a one dimension value, t, in the range [0, 7^Order-1] to 
// axial coordinates (q, r) on the hexagonal grid.
func (g *Gosper) Map(t int) (q, r int, err error) {
	if t < 0 || t >= g.N {
		return -1, -1, ErrOutOfRange
	}
	return g.Path[t][0], g.Path[t][1], nil
}

// MapInverse transform coordinates on Gosper curve from (q,r) to t.
func (g *Gosper) MapInverse(q, r int) (t int, err error) {
	for i, p := range g.Path {
		if p[0] == q && p[1] == r {
			return i, nil
		}
	}
	return -1, ErrOutOfRange
}
