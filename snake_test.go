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

func TestSnake(t *testing.T) {
	testCurve(t, func(n int) (SpaceFilling2D, error) {
		s, err := NewSnake(n)
		if err == nil {
			return s, nil
		}
		return nil, err
	}, GridSquare, 16, 0, ErrNotPositive, []curveTestCase{
		{0, 0, 0},
		{1, 1, 0},
		{15, 15, 0},
		{16, 15, 1},
		{17, 14, 1},
		{255, 0, 15},
	})
}
