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
	"time"
)

func main() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	// 重置定时器
	ticker.Reset(1 * time.Second)
	second := time.Now().Second()
	for t := range ticker.C {
		// 1s
		fmt.Printf("周期：%d 秒", t.Second()-second)
		break
	}
}
