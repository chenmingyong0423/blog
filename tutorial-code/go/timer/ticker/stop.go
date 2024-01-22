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
	ticker := time.NewTicker(time.Second)
	quit := make(chan struct{}) // 创建一个退出通道

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("定时器触发啦！")
			case <-quit:
				fmt.Println("协程停止啦！")
				return // 接收到退出信号，退出循环
			}
		}
	}()

	time.Sleep(time.Second * 3)
	ticker.Stop() // 停止定时器
	close(quit)   // 发送退出信号
	fmt.Println("定时器停止啦！")
}
