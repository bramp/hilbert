// Copyright 2015 Google Inc. All Rights Reserved.
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

import "errors"

// Errors returned when validating input.
var (
	ErrNotPositive     = errors.New("n must be greater than zero")
	ErrNotPowerOfTwo   = errors.New("n must be a power of two")
	ErrNotPowerOfThree = errors.New("n must be a power of three")
	ErrOutOfRange      = errors.New("value is out of range")
)

// GridType represents the underlying geometry of the 2D space.
type GridType int

const (
	GridSquare     GridType = iota // (x,y) Cartesian coordinates
	GridHexagonal                  // (q,r) Axial coordinates
	GridTriangular                 // (u,v) Triangular coordinates
)

// SpaceFilling2D represents a 2D space-filling curve that can map points from one dimensions to two.
type SpaceFilling2D interface {
	// Map transforms a one dimension value, t, in the range [0, n^2-1] to coordinates on the
	// curve in the two-dimension space.
	Map(t int) (x, y int, err error)

	// MapInverse transform coordinates on the curve from (x,y) to t.
	MapInverse(x, y int) (t int, err error)

	// GetDimensions returns the width and height of the 2D space.
	GetDimensions() (x, y int)

	// GetCount returns the total number of points on the curve.
	GetCount() int

	// GetGridType returns the geometry of the grid.
	GetGridType() GridType
}

// SpaceFilling is an alias for SpaceFilling2D for backward compatibility.
// Deprecated: Use SpaceFilling2D instead.
type SpaceFilling = SpaceFilling2D

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}
