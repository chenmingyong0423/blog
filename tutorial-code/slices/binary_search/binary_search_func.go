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
		{"Gopher", 24},
		{"Harry", 18},
	}

	idx, b := slices.BinarySearchFunc(users, User{Name: "Gopher"}, func(src User, dst User) int {
		return cmp.Compare(src.Name, dst.Name)
	})
	fmt.Println("Gopher:", idx, b)
}
