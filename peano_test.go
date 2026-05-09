// Copyright 2016 Google Inc. All Rights Reserved.
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
	"math/rand"
	"testing"
)

const peanoBenchmarkN = 81

func TestPeano(t *testing.T) {
	testCurve(t, func(n int) (SpaceFilling, error) {
		s, err := NewPeano(n)
		if err == nil {
			return s, nil
		}
		return nil, err
	}, 9, 2, ErrNotPowerOfThree, []curveTestCase{
		{0, 0, 0},
		{1, 0, 1},
		{2, 0, 2},
		{3, 1, 2},
		{4, 1, 1},
		{5, 1, 0},
		{6, 2, 0},
		{7, 2, 1},
		{8, 2, 2},
		{9, 2, 3},
		{80, 8, 8},
	})
}

func BenchmarkPeanoMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, err := NewPeano(peanoBenchmarkN)
		if err != nil {
			b.Fatalf("NewPeano(%d) failed: %s", peanoBenchmarkN, err)
		}
		for d := 0; d < peanoBenchmarkN*peanoBenchmarkN; d++ {
			s.Map(d)
		}
	}
}

func BenchmarkPeanoMapRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, err := NewPeano(peanoBenchmarkN)
		if err != nil {
			b.Fatalf("NewPeano(%d) failed: %s", peanoBenchmarkN, err)
		}
		for d := 0; d < peanoBenchmarkN*peanoBenchmarkN; d++ {
			rd := rand.Intn(peanoBenchmarkN * peanoBenchmarkN) // Pick a random d
			s.Map(rd)
		}
	}
}

func BenchmarkPeanoMapInverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, err := NewPeano(peanoBenchmarkN)
		if err != nil {
			b.Fatalf("Failed to create peano space: %s", err)
		}

		for x := 0; x < peanoBenchmarkN; x++ {
			for y := 0; y < peanoBenchmarkN; y++ {
				s.MapInverse(x, y)
			}
		}
	}
}

func TestIsPow3(t *testing.T) {
	testCases := []struct {
		in   float64
		want bool
	}{
		{-1, false},
		{0, false},
		{1, true},
		{2, false},
		{3, true},
		{3.1, false},
		{4, false},
		{5, false},
		{8.9999, false},
		{9, true},
		{9.00001, false},
		{27, true},
		{59049, true},
	}

	for _, tc := range testCases {
		got := isPow3(tc.in)
		if got != tc.want {
			t.Errorf("isPow3(%f) = %t want %t", tc.in, got, tc.want)
		}
	}
}
