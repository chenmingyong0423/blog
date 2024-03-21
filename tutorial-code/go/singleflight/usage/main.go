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
	"errors"
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync"
)

var errRedisKeyNotFound = errors.New("redis: key not found")

func fetchDataFromCache() (any, error) {
	fmt.Println("fetch data from cache")
	return nil, errRedisKeyNotFound
}

func fetchDataFromDataBase() (any, error) {
	fmt.Println("fetch data from database")
	return "程序员陈明勇", nil
}

func fetchData() (any, error) {
	cache, err := fetchDataFromCache()
	if err != nil && errors.Is(err, errRedisKeyNotFound) {
		fmt.Println(errRedisKeyNotFound.Error())
		return fetchDataFromDataBase()
	}
	return cache, err
}

func main() {
	var (
		sg singleflight.Group
		wg sync.WaitGroup
	)

	for range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			v, err, shared := sg.Do("key", fetchData)
			if err != nil {
				panic(err)
			}
			fmt.Printf("v: %v, shared: %v\n", v, shared)
		}()
	}
	wg.Wait()
}
