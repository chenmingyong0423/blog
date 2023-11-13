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
	"chenmingyong0423/blog/tutorial-code/go-mongox/collection"
	"context"
	"fmt"

	"github.com/chenmingyong0423/go-mongox"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// 你需要预先创建一个 *mongo.Collection 对象
	mongoCollection := collection.NewCollection()
	// 使用 Post 结构体作为泛型参数创建一个 collection
	postCollection := mongox.NewCollection[collection.Post](mongoCollection)

	docs := []collection.Post{
		{Id: "1", Title: "go", Author: "陈明勇", Content: "..."},
		{Id: "2", Title: "mongo", Author: "陈明勇", Content: "..."},
	}
	manyResult, err := postCollection.Creator().InsertMany(context.Background(), docs)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(manyResult.InsertedIDs) == 2) // true
	// 携带 option 参数
	manyResult, err = postCollection.Creator().ManyOptions(options.InsertMany().SetComment("test")).InsertMany(context.Background(), []collection.Post{
		{Id: "3", Title: "go-mongox", Author: "陈明勇", Content: "..."},
		{Id: "4", Title: "builder", Author: "陈明勇", Content: "..."},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(len(manyResult.InsertedIDs) == 2) // true
}
