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
	"context"
	"github.com/chenmingyong0423/go-mongox"
	"github.com/chenmingyong0423/go-mongox/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

// Sample code is not the best way to create it
func newCollection() *mongo.Collection {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
		Username:   "test",
		Password:   "test",
		AuthSource: "db-test",
	}))
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		panic(err)
	}
	collection := client.Database("db-test").Collection("test_post")
	return collection
}

type Post struct {
	Id      string `bson:"_id"`
	Title   string `bson:"title"`
	Author  string `bson:"author"`
	Content string `bson:"content"`
}

func main() {
	mongoCollection := newCollection()

	initData(mongoCollection)
	defer removeData(mongoCollection)

	postCollection := mongox.NewCollection[Post](mongoCollection)

	// 通过 map 构造查询条件
	// - 根据 id 查询文档
	post, err := postCollection.Finder().Filter(map[string]any{"_id": "1"}).FindOne(context.Background())
	if err != nil {
		panic(err)
	}
	log.Default().Printf("%+v\n", post) // &{Id:1 Title:go 语言 go-mongox 库的使用教程 Author:陈明勇 Content:go-mongox 旨在提供更方便和高效的MongoDB数据操作体验。}

	// 通过结构体构造查询条件
	// - 根据 title 查询文档
	type PostTitle struct {
		Title string `bson:"title"`
	}
	post2, err := postCollection.Finder().Filter(PostTitle{Title: "go 语言 go-mongox 库的使用教程"}).FindOne(context.Background())
	if err != nil {
		panic(err)
	}
	log.Default().Printf("%+v\n", post2) // &{Id:1 Title:go 语言 go-mongox 库的使用教程 Author:陈明勇 Content:go-mongox 旨在提供更方便和高效的MongoDB数据操作体验。}

	// 通过 KeyValue 构造查询条件
	// - 根据 author 和 title 查询文档
	post3, err := postCollection.Finder().FilterKeyValue(converter.KeyValue("title", "go 语言 go-mongox 库的使用教程"), converter.KeyValue("author", "陈明勇")).FindOne(context.Background())
	if err != nil {
		panic(err)
	}
	log.Default().Printf("%+v\n", post3) // &{Id:1 Title:go 语言 go-mongox 库的使用教程 Author:陈明勇 Content:go-mongox 旨在提供更方便和高效的MongoDB数据操作体验。}
}

// 预先插入一条数据，方便测试
func initData(mongoCollection *mongo.Collection) {
	_, err := mongoCollection.InsertOne(context.Background(), Post{Id: "1", Title: "go 语言 go-mongox 库的使用教程", Author: "陈明勇", Content: "go-mongox 旨在提供更方便和高效的MongoDB数据操作体验。"})
	if err != nil {
		panic(err)
	}
}

func removeData(mongoCollection *mongo.Collection) {
	_, err := mongoCollection.DeleteOne(context.Background(), bson.D{{Key: "_id", Value: "1"}})
	if err != nil {
		panic(err)
	}
}
