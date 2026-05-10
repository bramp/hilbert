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

import (
	"math/rand"
	"testing"
)

const benchmarkN = 32

func TestHilbert(t *testing.T) {
	testCurve(t, func(n int) (SpaceFilling2D, error) {
		s, err := NewHilbert(n)
		if err == nil {
			return s, nil
		}
		return nil, err
	}, GridSquare, 16, 3, ErrNotPowerOfTwo, []curveTestCase{
		{0, 0, 0},
		{16, 4, 0},
		{32, 4, 4},
		{48, 3, 7},
		{64, 0, 8},
		{80, 0, 12},
		{96, 4, 12},
		{112, 7, 11},
		{128, 8, 8},
		{144, 8, 12},
		{160, 12, 12},
		{170, 15, 15},
		{176, 15, 11},
		{192, 15, 7},
		{208, 11, 7},
		{224, 11, 3},
		{240, 12, 0},
		{255, 15, 0},
	})
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, err := NewHilbert(benchmarkN)
		if err != nil {
			b.Fatalf("Failed to create hibert space: %s", err)
		}
		for d := 0; d < benchmarkN*benchmarkN; d++ {
			s.Map(d)
		}
	}
}

func BenchmarkMapRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, err := NewHilbert(benchmarkN)
		if err != nil {
			b.Fatalf("Failed to create hibert space: %s", err)
		}
		for d := 0; d < benchmarkN*benchmarkN; d++ {
			rd := rand.Intn(benchmarkN * benchmarkN) // Pick a random d
			s.Map(rd)
		}
	}
}

func BenchmarkMapInverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, err := NewHilbert(benchmarkN)
		if err != nil {
			b.Fatalf("Failed to create hibert space: %s", err)
		}

		for x := 0; x < benchmarkN; x++ {
			for y := 0; y < benchmarkN; y++ {
				s.MapInverse(x, y)
			}
		}
	}
}
