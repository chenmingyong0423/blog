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
	// 使用 NewTimer 创建一个定时器
	timer := time.NewTimer(time.Second)
	go func() {
		select {
		case <-timer.C:
			fmt.Println("timer 定时器触发啦！")
		}
	}()
	// 使用 AfterFunc 创建另一个定时器
	time.AfterFunc(time.Second, func() {
		fmt.Println("timer2 定时器触发啦！")
	})
	// 主goroutine等待两秒，确保看到定时器触发的输出
	time.Sleep(time.Second * 2)
}
