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
	"github.com/chenmingyong0423/go-mongox/bsonx"
	"github.com/chenmingyong0423/go-mongox/builder/query"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()

	// 你需要预先创建一个 *mongo.Collection 对象
	mongoCollection := collection.NewCollection()
	// 使用 Post 结构体作为泛型参数创建一个 collection
	postCollection := mongox.NewCollection[collection.Post](mongoCollection)
	_, err := postCollection.Creator().InsertMany(ctx, []collection.Post{
		{Title: "go-mongox", Author: "陈明勇", Content: "go-mongox 旨在提供更方便和高效的MongoDB数据操作体验。"},
		{Title: "mongo", Author: "陈明勇", Content: "..."},
	})
	if err != nil {
		panic(err)
	}

	// 查询多个文档
	// bson.D{bson.E{Key: "title", Value: bson.M{"$in": []string{"go-mongox", "mongo"}}}}
	posts, err := postCollection.Finder().
		Filter(query.In("title", []string{"go-mongox", "mongo"}...)).
		Find(ctx)
	if err != nil {
		panic(err)
	}
	for _, post := range posts {
		fmt.Println(post)
	}

	// 设置 *options.FindOptions 参数
	// bson.D{bson.E{Key: "title", Value: bson.M{"$in": []string{"go-mongox", "mongo"}}}}
	posts2, err := postCollection.Finder().
		Filter(query.In("title", []string{"go-mongox", "mongo"}...)).
		Find(ctx, options.Find().SetProjection(bsonx.M("content", 0)))
	if err != nil {
		panic(err)
	}
	for _, post := range posts2 {
		fmt.Println(post)
	}
}
