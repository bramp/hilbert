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

func TestMooreNewErrors(t *testing.T) {
	var newTestCases = []struct {
		n    int
		want error
	}{
		{-1, ErrNotPositive},
		{0, ErrNotPositive},
		{3, ErrNotPowerOfTwo},
	}

	for _, tc := range newTestCases {
		s, err := NewMoore(tc.n)
		if s != nil || err != tc.want {
			t.Errorf("NewMoore(%d) = (%+v, %q) did not fail want (?, %q)", tc.n, s, err, tc.want)
		}
	}
}

func TestMooreMapRangeErrors(t *testing.T) {
	var mapRangeTestCases = []struct {
		t       int
		wantErr error
	}{
		{-1, ErrOutOfRange},
		{0, nil},
		{15, nil},
		{16, ErrOutOfRange},
	}

	s, err := NewMoore(4)
	if err != nil {
		t.Fatalf("NewMoore(4) failed: %s", err)
	}

	for _, tc := range mapRangeTestCases {
		if _, _, err = s.Map(tc.t); err != tc.wantErr {
			t.Errorf("Map(%d) = %q want %q", tc.t, err, tc.wantErr)
		}
	}
}

func TestMooreMapInverseRangeErrors(t *testing.T) {
	var mapInverseRangeTestCases = []struct {
		x, y    int
		wantErr error
	}{
		{0, 0, nil},
		{3, 3, nil},
		{-1, 0, ErrOutOfRange},
		{0, -1, ErrOutOfRange},
		{4, 0, ErrOutOfRange},
		{0, 4, ErrOutOfRange},
	}

	s, err := NewMoore(4)
	if err != nil {
		t.Fatalf("NewMoore(4) failed: %s", err)
	}

	for _, tc := range mapInverseRangeTestCases {
		if _, err = s.MapInverse(tc.x, tc.y); err != tc.wantErr {
			t.Errorf("MapInverse(%d, %d) = %q want %q", tc.x, tc.y, err, tc.wantErr)
		}
	}
}

func TestMooreMap(t *testing.T) {
	s, err := NewMoore(4)
	if err != nil {
		t.Fatalf("NewMoore(4) failed: %s", err)
	}

	// Test a few specific values
	// For n=4, m2=2.
	// t=0: q=0, tSub=0. hSub(2).Map(0)=(0,0). Moore: (2-1-0, 0) = (1, 0)
	// t=1: q=0, tSub=1. hSub(2).Map(1)=(0,1). Moore: (2-1-1, 0) = (0, 0)
	var testCases = []struct {
		t, x, y int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, 0, 1},
		{3, 1, 1},
		{4, 1, 2},
		{5, 0, 2},
		{6, 0, 3},
		{7, 1, 3},
		{8, 2, 3},
		{9, 3, 3},
		{10, 3, 2},
		{11, 2, 2},
		{12, 2, 1},
		{13, 3, 1},
		{14, 3, 0},
		{15, 2, 0},
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

func TestMooreAllMapValues(t *testing.T) {
	s, err := NewMoore(16)
	if err != nil {
		t.Fatalf("NewMoore(16) failed: %s", err)
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
