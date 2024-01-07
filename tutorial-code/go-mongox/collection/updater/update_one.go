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

	"github.com/chenmingyong0423/go-mongox/builder/query"

	"github.com/chenmingyong0423/go-mongox"
	"github.com/chenmingyong0423/go-mongox/bsonx"
	"github.com/chenmingyong0423/go-mongox/builder/update"
	"github.com/chenmingyong0423/go-mongox/types"
)

func main() {
	ctx := context.Background()

	// 你需要预先创建一个 *mongo.Collection 对象
	mongoCollection := collection.NewCollection()
	// 使用 Post 结构体作为泛型参数创建一个 collection
	postCollection := mongox.NewCollection[collection.Post](mongoCollection)
	oneResult, err := postCollection.Creator().InsertOne(ctx, collection.Post{Title: "go-mongox", Author: "陈明勇", Content: "go-mongox 旨在提供更方便和高效的MongoDB数据操作体验。"})
	if err != nil {
		panic(err)
	}

	// 更新单个文档
	// 通过 update 包构建 bson 更新语句
	updateResult, err := postCollection.Updater().
		Filter(query.Id(oneResult.InsertedID)).
		Updates(update.Set("title", "golang")).
		UpdateOne(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(updateResult.ModifiedCount == 1) // true

	// - 使用 map 构造更新数据，并设置 *options.UpdateOptions，执行 upsert 操作
	updateResult2, err := postCollection.Updater().
		Filter(bsonx.Id("custom-id")).
		UpdatesWithOperator(types.Set, bsonx.M("title", "mongo")).
		UpdateOne(ctx, options.Update().SetUpsert(true))
	if err != nil {
		panic(err)
	}
	fmt.Println(updateResult2.UpsertedID.(string) == "custom-id") // true
}
