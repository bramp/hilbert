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

// Moore represents a 2D Moore curve of order N for mapping to and from.
// The Moore curve is a closed-loop variation of the Hilbert curve.
// Implements SpaceFilling interface.
type Moore struct {
	Hilbert
}

// NewMoore returns a Moore space which maps integers to and from the curve.
// n must be a power of two.
func NewMoore(n int) (*Moore, error) {
	h, err := NewHilbert(n)
	if err != nil {
		return nil, err
	}
	return &Moore{Hilbert: *h}, nil
}

// Map transforms a one dimension value, t, in the range [0, n^2-1] to coordinates on the Moore
// curve in the two-dimension space, where x and y are within [0,n-1].
func (m *Moore) Map(t int) (x, y int, err error) {
	if t < 0 || t >= m.N*m.N {
		return -1, -1, ErrOutOfRange
	}

	if m.N == 1 {
		return 0, 0, nil
	}

	m2 := m.N / 2
	m2sq := m2 * m2
	q := t / m2sq
	tSub := t % m2sq

	hSub, _ := NewHilbert(m2)
	xh, yh, _ := hSub.Map(tSub)

	switch q {
	case 0: // Bottom-Left
		return m2 - 1 - yh, xh, nil
	case 1: // Top-Left
		return m2 - 1 - yh, xh + m2, nil
	case 2: // Top-Right
		return yh + m2, m.N - 1 - xh, nil
	case 3: // Bottom-Right
		return yh + m2, m2 - 1 - xh, nil
	}

	return -1, -1, ErrOutOfRange
}

// MapInverse transform coordinates on Moore curve from (x,y) to t.
func (m *Moore) MapInverse(x, y int) (t int, err error) {
	if x < 0 || x >= m.N || y < 0 || y >= m.N {
		return -1, ErrOutOfRange
	}

	if m.N == 1 {
		return 0, nil
	}

	m2 := m.N / 2
	m2sq := m2 * m2
	var q int
	var xh, yh int

	if x < m2 && y < m2 { // Quadrant 0
		q = 0
		xh = y
		yh = m2 - 1 - x
	} else if x < m2 && y >= m2 { // Quadrant 1
		q = 1
		xh = y - m2
		yh = m2 - 1 - x
	} else if x >= m2 && y >= m2 { // Quadrant 2
		q = 2
		xh = m.N - 1 - y
		yh = x - m2
	} else { // Quadrant 3
		q = 3
		xh = m2 - 1 - y
		yh = x - m2
	}

	hSub, _ := NewHilbert(m2)
	tSub, _ := hSub.MapInverse(xh, yh)

	return q*m2sq + tSub, nil
}

