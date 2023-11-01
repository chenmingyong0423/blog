// Copyright 2023 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	ints := []int{1, 2, 5, 3, 4}
	slices.Sort(ints)
	floats := []float64{2.0, 3.0, math.NaN(), 1.0}
	slices.Sort(floats)
	fmt.Println(ints)
	fmt.Println(floats)
}
