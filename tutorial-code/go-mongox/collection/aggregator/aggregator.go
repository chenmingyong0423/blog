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
	"github.com/chenmingyong0423/go-mongox/builder/aggregation"
	"github.com/chenmingyong0423/go-mongox/types"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	// 你需要预先创建一个 *mongo.Collection 对象
	mongoCollection := collection.NewCollection()
	// 使用 Post 结构体作为泛型参数创建一个 collection
	postCollection := mongox.NewCollection[collection.Post](mongoCollection)
	_, err := postCollection.Creator().InsertMany(context.Background(), []collection.Post{
		{Id: "1", Title: "go", Author: "陈明勇", Content: "..."},
		{Id: "2", Title: "mongo", Author: "陈明勇", Content: "..."},
	})
	if err != nil {
		panic(err)
	}

	// 聚合操作
	// - 使用 aggregation 包构造 pipeline
	posts, err := postCollection.
		Aggregator().Pipeline(aggregation.StageBsonBuilder().Project(bsonx.M("content", 0)).Build()).
		Aggregate(context.Background())
	if err != nil {
		panic(err)
	}
	for _, post := range posts {
		fmt.Println(post)
	}

	// 如果我们通过聚合操作更改字段的名称，那么我们可以使用 AggregationWithCallback 方法，然后通过 callback 函数将结果映射到我们预期的结构体中
	type DiffPost struct {
		Id          string `bson:"_id"`
		Title       string `bson:"title"`
		Name        string `bson:"name"` // author → name
		Content     string `bson:"content"`
		Outstanding bool   `bson:"outstanding"`
	}
	result := make([]*DiffPost, 0)
	//  将 author 字段更名为 name，排除 content 字段，添加 outstanding 字段，返回结果为 []*DiffPost
	err = postCollection.Aggregator().
		Pipeline(aggregation.StageBsonBuilder().Project(
			bsonx.D(types.KV("name", "$author"), types.KV("author", 1), types.KV("_id", 1), types.KV("title", 1), types.KV("outstanding", aggregation.BsonBuilder().Eq("$author", "陈明勇").Build()))).Build(),
		).
		AggregateWithCallback(context.Background(), func(ctx context.Context, cursor *mongo.Cursor) error {
			return cursor.All(ctx, &result)
		})

	for _, post := range result {
		fmt.Println(post)
	}
}
