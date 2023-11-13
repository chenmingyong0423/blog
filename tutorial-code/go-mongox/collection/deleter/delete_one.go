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
	"go.mongodb.org/mongo-driver/mongo/options"
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

	// 删除单个文档
	deleteResult, err := postCollection.Deleter().Filter(bsonx.Id("1")).DeleteOne(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(deleteResult.DeletedCount == 1) // true

	// 携带 option 参数
	deleteResult2, err := postCollection.Deleter().Filter(bsonx.Id("2")).Options(options.Delete().SetComment("test")).DeleteOne(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(deleteResult2.DeletedCount == 1) // true
}
