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
	"sync"
)

func main() {
	var wg sync.WaitGroup
	workers := []string{"陈明勇", "Mingyong Chen", "chenmingyong"}
	wg.Add(3)
	for i, worker := range workers {
		go Work(worker, i+1, &wg)
	}
	// 等待所有任务执行完毕
	wg.Wait()
	fmt.Println("所有任务执行完毕！")
}

func Work(name string, no int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s 正在执行任务 %d...\n", name, no)
}
