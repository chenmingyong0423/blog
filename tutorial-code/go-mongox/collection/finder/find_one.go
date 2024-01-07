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
	"github.com/chenmingyong0423/go-mongox/bsonx"
	"github.com/chenmingyong0423/go-mongox/builder/query"
)

func main() {
	// 你需要预先创建一个 *mongo.Collection 对象
	mongoCollection := collection.NewCollection()
	// 使用 Post 结构体作为泛型参数创建一个 collection
	postCollection := mongox.NewCollection[collection.Post](mongoCollection)
	oneResult, err := postCollection.Creator().InsertOne(context.Background(), collection.Post{Title: "go-mongox", Author: "陈明勇", Content: "go-mongox 旨在提供更方便和高效的MongoDB数据操作体验。"})
	if err != nil {
		panic(err)
	}

	// 查询单个文档
	post, err := postCollection.Finder().Filter(query.Id(oneResult.InsertedID)).FindOne(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(post)

	// 设置 *options.FindOneOptions 参数
	post2, err := postCollection.Finder().
		Filter(query.Id(oneResult.InsertedID)).
		FindOne(context.Background(), options.FindOne().SetProjection(bsonx.M("content", 0)))
	if err != nil {
		panic(err)
	}
	fmt.Println(post2)

	// - map 作为 filter 条件
	post3, err := postCollection.Finder().Filter(map[string]any{"_id": oneResult.InsertedID}).FindOne(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(post3)

	// - 复杂条件查询
	// -- 使用 query 包构造复杂的 bson: bson.D{bson.E{Key: "title", Value: bson.M{"$eq": "go"}}, bson.E{Key: "author", Value: bson.M{"$eq": "陈明勇"}}}
	post4, err := postCollection.Finder().
		Filter(query.BsonBuilder().Eq("title", "go-mongox").Eq("author", "陈明勇").Build()).
		FindOne(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(post4)
}
