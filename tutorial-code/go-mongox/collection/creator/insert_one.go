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

	// 插入一个文档
	doc := collection.Post{Id: "1", Title: "go-mongox", Author: "陈明勇", Content: "go-mongox，不一样的体验。"}
	oneResult, err := postCollection.Creator().InsertOne(context.Background(), doc)
	if err != nil {
		panic(err)
	}
	fmt.Println(oneResult.InsertedID.(string) == "1") // true

	// 携带 option 参数
	oneResult, err = postCollection.Creator().OneOptions(options.InsertOne().SetComment("test")).InsertOne(context.Background(), collection.Post{Id: "2", Title: "go 语言 go-mongox 库的使用教程", Author: "陈明勇", Content: "go-mongox 旨在提供更方便和高效的MongoDB数据操作体验。"})
	if err != nil {
		panic(err)
	}
	fmt.Println(oneResult.InsertedID.(string) == "2") // true
}
