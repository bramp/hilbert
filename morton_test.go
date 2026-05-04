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
	"testing"
)

func TestMortonNewErrors(t *testing.T) {
	var newTestCases = []struct {
		n    int
		want error
	}{
		{-1, ErrNotPositive},
		{0, ErrNotPositive},
		{3, ErrNotPowerOfTwo},
	}

	for _, tc := range newTestCases {
		s, err := NewMorton(tc.n)
		if s != nil || err != tc.want {
			t.Errorf("NewMorton(%d) = (%+v, %q) did not fail want (?, %q)", tc.n, s, err, tc.want)
		}
	}
}

func TestMortonMapRangeErrors(t *testing.T) {
	var mapRangeTestCases = []struct {
		t       int
		wantErr error
	}{
		{-1, ErrOutOfRange},
		{0, nil},
		{15, nil},
		{16, ErrOutOfRange},
	}

	s, err := NewMorton(4)
	if err != nil {
		t.Fatalf("NewMorton(4) failed: %s", err)
	}

	for _, tc := range mapRangeTestCases {
		if _, _, err = s.Map(tc.t); err != tc.wantErr {
			t.Errorf("Map(%d) = %q want %q", tc.t, err, tc.wantErr)
		}
	}
}

func TestMortonMapInverseRangeErrors(t *testing.T) {
	var mapInverseRangeTestCases = []struct {
		x, y    int
		wantErr error
	} {
		{0, 0, nil},
		{3, 3, nil},
		{-1, 0, ErrOutOfRange},
		{0, -1, ErrOutOfRange},
		{4, 0, ErrOutOfRange},
		{0, 4, ErrOutOfRange},
	}

	s, err := NewMorton(4)
	if err != nil {
		t.Fatalf("NewMorton(4) failed: %s", err)
	}

	for _, tc := range mapInverseRangeTestCases {
		if _, err = s.MapInverse(tc.x, tc.y); err != tc.wantErr {
			t.Errorf("MapInverse(%d, %d) = %q want %q", tc.x, tc.y, err, tc.wantErr)
		}
	}
}

func TestMortonMap(t *testing.T) {
	s, err := NewMorton(4)
	if err != nil {
		t.Fatalf("NewMorton(4) failed: %s", err)
	}

	var testCases = []struct {
		t, x, y int
	}{
		{0, 0, 0},
		{1, 0, 1},
		{2, 1, 0},
		{3, 1, 1},
		{4, 0, 2},
		{5, 0, 3},
		{6, 1, 2},
		{7, 1, 3},
		{8, 2, 0},
		{15, 3, 3},
	}

	for _, tc := range testCases {
		x, y, err := s.Map(tc.t)
		if err != nil {
			t.Errorf("Map(%d) returned error: %s", tc.t, err)
		}
		if x != tc.x || y != tc.y {
			t.Errorf("Map(%d) = (%d, %d) want (%d, %d)", tc.t, x, y, tc.x, tc.y)
		}

		tPrime, err := s.MapInverse(x, y)
		if err != nil {
			t.Errorf("MapInverse(%d, %d) returned error: %s", x, y, err)
		}
		if tPrime != tc.t {
			t.Errorf("MapInverse(%d, %d) = %d want %d", x, y, tPrime, tc.t)
		}
	}
}

func TestMortonAllMapValues(t *testing.T) {
	s, err := NewMorton(16)
	if err != nil {
		t.Fatalf("NewMorton(16) failed: %s", err)
	}

	for tVal := 0; tVal < s.N*s.N; tVal++ {
		x, y, err := s.Map(tVal)
		if err != nil {
			t.Errorf("Map(%d) returned error: %s", tVal, err)
		}
		if x < 0 || x >= s.N || y < 0 || y >= s.N {
			t.Errorf("Map(%d) = (%d, %d) out of range", tVal, x, y)
		}

		tPrime, err := s.MapInverse(x, y)
		if err != nil {
			t.Errorf("MapInverse(%d, %d) returned error: %s", x, y, err)
		}
		if tPrime != tVal {
			t.Errorf("Map(%d) -> MapInverse(%d, %d) = %d want %d", tVal, x, y, tPrime, tVal)
		}
	}
}
