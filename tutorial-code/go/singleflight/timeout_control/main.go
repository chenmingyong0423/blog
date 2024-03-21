// Copyright 2024 chenmingyong0423

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
	"golang.org/x/sync/singleflight"
	"time"
)

func main() {
	var sg singleflight.Group
	doChan := sg.DoChan("key", func() (interface{}, error) {
		time.Sleep(4 * time.Second)
		return "程序员陈明勇", nil
	})
	select {
	case <-doChan:
		fmt.Println("done")
	case <-time.After(2 * time.Second):
		fmt.Println("timeout")
		// 采用其他降级策略
	}
}
