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

	"github.com/chenmingyong0423/go-eventbus"
)

func main() {
	eventBus := eventbus.NewEventBus()

	// 订阅 post 主题事件
	subscribe := eventBus.Subscribe("post")

	go func() {
		for event := range subscribe {
			fmt.Println(event.Payload)

		}
	}()

	eventBus.Publish("post", eventbus.Event{Payload: map[string]any{
		"postId": 1,
		"title":  "Go 事件驱动编程：实现一个简单的事件总线",
		"author": "陈明勇",
	}})
	// 不存在订阅者的 topic
	eventBus.Publish("pay", eventbus.Event{Payload: "pay"})

	time.Sleep(time.Second * 2)
	// 取消订阅 post 主题事件
	eventBus.Unsubscribe("post", subscribe)
}
