// Copyright 2026 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance_ with the License.
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

func TestMoore(t *testing.T) {
	testCurve(t, func(n int) (SpaceFilling2D, error) {
		s, err := NewMoore(n)
		if err == nil {
			return s, nil
		}
		return nil, err
	}, GridSquare, 16, 3, ErrNotPowerOfTwo, []curveTestCase{
		{0, 7, 0},
		{1, 6, 0},
		{2, 6, 1},
		{3, 7, 1},
	})
}
