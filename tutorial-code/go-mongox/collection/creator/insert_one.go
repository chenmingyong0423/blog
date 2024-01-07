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

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/chenmingyong0423/go-mongox"
)

func main() {
	// 你需要预先创建一个 *mongo.Collection 对象
	mongoCollection := collection.NewCollection()
	// 使用 Post 结构体作为泛型参数创建一个 collection
	postCollection := mongox.NewCollection[collection.Post](mongoCollection)

	ctx := context.Background()
	// 插入一个文档
	doc := collection.Post{Title: "go 语言 go-mongox 库的使用教程", Author: "陈明勇", Content: "go-mongox 旨在提供更方便和高效的MongoDB数据操作体验。"}
	oneResult, err := postCollection.Creator().InsertOne(ctx, doc)
	if err != nil {
		panic(err)
	}
	fmt.Println(oneResult.InsertedID != nil) // true

	// 携带 option 参数
	oneResult, err = postCollection.Creator().InsertOne(ctx, doc, options.InsertOne().SetComment("test"))
	if err != nil {
		panic(err)
	}
	fmt.Println(oneResult.InsertedID != nil) // true
}
