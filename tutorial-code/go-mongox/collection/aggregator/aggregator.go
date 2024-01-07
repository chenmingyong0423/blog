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
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	ctx := context.Background()

	// 你需要预先创建一个 *mongo.Collection 对象
	mongoCollection := collection.NewCollection()
	// 使用 Post 结构体作为泛型参数创建一个 collection
	postCollection := mongox.NewCollection[collection.Post](mongoCollection)
	_, err := postCollection.Creator().InsertMany(ctx, []collection.Post{
		{Title: "go-mongox", Author: "陈明勇", Content: "go-mongox 旨在提供更方便和高效的MongoDB数据操作体验。"},
		{Title: "builder", Author: "陈明勇", Content: "..."},
	})
	if err != nil {
		panic(err)
	}

	// 聚合操作
	// - 使用 aggregation 包构造 pipeline
	posts, err := postCollection.
		// 去除 content 字段
		Aggregator().Pipeline(aggregation.StageBsonBuilder().Project(bsonx.M("content", 0)).Build()).
		Aggregate(ctx)
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
		Pipeline(
			aggregation.StageBsonBuilder().Project(
				bsonx.NewD().Add("name", "$author").Add("author", 1).Add("_id", 1).Add("title", 1).Add("outstanding", aggregation.Eq("$author", "陈明勇")).Build(),
			).Build(),
		).
		AggregateWithCallback(ctx, func(ctx context.Context, cursor *mongo.Cursor) error {
			return cursor.All(ctx, &result)
		})

	for _, post := range result {
		fmt.Println(post)
	}
}
