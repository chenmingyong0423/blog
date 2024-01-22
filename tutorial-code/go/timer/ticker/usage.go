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
	"context"
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()

	go func() {
		for {
			select {
			case <-timeout.Done():
				fmt.Println("timeout done")
				return
			case <-ticker.C:
				fmt.Println("定时器触发啦！")
			}
		}
	}()

	// 主goroutine等待 7 秒，确保看到定时器触发的输出
	time.Sleep(time.Second * 7)
}
