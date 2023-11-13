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
	"github.com/chenmingyong0423/go-mongox/builder/update"
	"github.com/chenmingyong0423/go-mongox/types"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// 你需要预先创建一个 *mongo.Collection 对象
	mongoCollection := collection.NewCollection()
	// 使用 Post 结构体作为泛型参数创建一个 collection
	postCollection := mongox.NewCollection[collection.Post](mongoCollection)
	_, err := postCollection.Creator().InsertOne(context.Background(), collection.Post{Id: "1", Title: "go", Author: "陈明勇", Content: "..."})
	if err != nil {
		panic(err)
	}

	// 更新单个文档
	// 通过 update 包构建 bson 更新语句
	updateResult, err := postCollection.Updater().
		Filter(bsonx.Id("1")).
		Updates(update.BsonBuilder().Set(bsonx.M("title", "golang")).Build()).
		UpdateOne(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(updateResult.ModifiedCount == 1) // true

	// - 使用 map 构造更新数据，并设置 *options.UpdateOptions，执行 upsert 操作
	updateResult2, err := postCollection.Updater().
		Filter(bsonx.Id("2")).
		UpdatesWithOperator(types.Set, map[string]any{"title": "mongo"}).Options(options.Update().SetUpsert(true)).
		UpdateOne(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(updateResult2.UpsertedID.(string) == "2") // true
}
