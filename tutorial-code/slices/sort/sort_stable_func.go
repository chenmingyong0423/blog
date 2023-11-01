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
	"cmp"
	"fmt"
	"slices"
)

func main() {
	type User struct {
		Name string
		Age  int
	}

	users := []User{
		{"Aaron", 20},
		{"Gopher", 16},
		{"Harry", 16},
		{"Burt", 18},
	}
	slices.SortStableFunc(users, func(a, b User) int {
		return cmp.Compare(a.Age, b.Age)
	})
	fmt.Println(users)
}
