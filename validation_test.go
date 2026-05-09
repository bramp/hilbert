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
	"fmt"
	"testing"
)

// curveTestCase represents a single mapping test case.
type curveTestCase struct {
	d, x, y int
}

// testCurve runs a standard battery of tests on a space filling curve implementation.
func testCurve(t *testing.T, constructor func(n int) (SpaceFilling, error), validN int, invalidN int, errN error, testCases []curveTestCase) {
	t.Helper()

	t.Run("NewErrors", func(t *testing.T) {
		s, err := constructor(invalidN)
		if s != nil {
			t.Errorf("New(%d) returned non-nil curve", invalidN)
		}
		if err == nil {
			t.Errorf("New(%d) returned nil error", invalidN)
		} else if err.Error() != errN.Error() {
			t.Errorf("New(%d) error = %q; want %q", invalidN, err.Error(), errN.Error())
		}

		s, err = constructor(-1)
		if s != nil {
			t.Errorf("New(-1) returned non-nil curve")
		}
		if err == nil {
			t.Errorf("New(-1) returned nil error")
		} else if err.Error() != ErrNotPositive.Error() {
			t.Errorf("New(-1) error = %q; want %q", err.Error(), ErrNotPositive.Error())
		}
	})

	s, err := constructor(validN)
	if err != nil {
		t.Fatalf("Failed to create curve with N=%d: %v", validN, err)
	}

	w, h := s.GetDimensions()
	if w != validN || h != validN {
		t.Errorf("GetDimensions() = (%d, %d); want (%d, %d)", w, h, validN, validN)
	}

	t.Run("RangeErrors", func(t *testing.T) {
		n := w * h
		if _, _, err := s.Map(-1); err != ErrOutOfRange {
			t.Errorf("Map(-1) error = %v; want %v", err, ErrOutOfRange)
		}
		if _, _, err := s.Map(n); err != ErrOutOfRange {
			t.Errorf("Map(%d) error = %v; want %v", n, err, ErrOutOfRange)
		}
		if _, err := s.MapInverse(-1, 0); err != ErrOutOfRange {
			t.Errorf("MapInverse(-1, 0) error = %v; want %v", err, ErrOutOfRange)
		}
		if _, err := s.MapInverse(0, h); err != ErrOutOfRange {
			t.Errorf("MapInverse(0, %d) error = %v; want %v", h, err, ErrOutOfRange)
		}
	})

	t.Run("SmallMap", func(t *testing.T) {
		s1, err := constructor(1)
		if err != nil {
			t.Fatalf("New(1) failed: %v", err)
		}
		x, y, err := s1.Map(0)
		if err != nil || x != 0 || y != 0 {
			t.Errorf("Map(0) = (%d, %d, %v); want (0, 0, nil)", x, y, err)
		}
		d, err := s1.MapInverse(0, 0)
		if err != nil || d != 0 {
			t.Errorf("MapInverse(0, 0) = (%d, %v); want (0, nil)", d, err)
		}
	})

	if len(testCases) > 0 {
		t.Run("SpecificCases", func(t *testing.T) {
			for _, tc := range testCases {
				x, y, err := s.Map(tc.d)
				if err != nil {
					t.Errorf("Map(%d) error: %v", tc.d, err)
					continue
				}
				if x != tc.x || y != tc.y {
					t.Errorf("Map(%d) = (%d, %d); want (%d, %d)", tc.d, x, y, tc.x, tc.y)
				}

				d, err := s.MapInverse(tc.x, tc.y)
				if err != nil {
					t.Errorf("MapInverse(%d, %d) error: %v", tc.x, tc.y, err)
					continue
				}
				if d != tc.d {
					t.Errorf("MapInverse(%d, %d) = %d; want %d", tc.x, tc.y, d, tc.d)
				}
			}
		})
	}

	t.Run("FillingAndRoundTrip", func(t *testing.T) {
		n := w * h
		visited := make(map[string]int)

		for i := 0; i < n; i++ {
			x, y, err := s.Map(i)
			if err != nil {
				t.Errorf("Map(%d) error: %v", i, err)
				continue
			}

			key := fmt.Sprintf("%d,%d", x, y)
			if firstI, exists := visited[key]; exists {
				t.Errorf("Map(%d) returned (%d, %d) already visited by Map(%d)", i, x, y, firstI)
			}
			visited[key] = i

			d, err := s.MapInverse(x, y)
			if err != nil {
				t.Errorf("MapInverse(%d, %d) error: %v", x, y, err)
			} else if d != i {
				t.Errorf("MapInverse(Map(%d)) = %d; want %d", i, d, i)
			}
		}

		if _, ok := s.(*Sierpinski); !ok {
			if len(visited) != n {
				t.Errorf("Failed to visit all squares. Visited %d/%d", len(visited), n)
			}
		}
	})
}
