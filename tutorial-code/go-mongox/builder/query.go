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
	"fmt"

	"github.com/chenmingyong0423/go-mongox/bsonx"
	"github.com/chenmingyong0423/go-mongox/builder/query"
)

func main() {
	// bson.D{bson.E{Key:"姓名", Value:"陈明勇"}}
	d := query.BsonBuilder().Add(bsonx.KV("姓名", "陈明勇")).Build()
	fmt.Printf("%#v\n\n", d)

	// bson.D{bson.E{Key:"age", Value:bson.D{{Key:"$gt", Value:18}, bson.E{Key:"$lt", Value:25}}}}
	d = query.BsonBuilder().Gt("age", 18).Lt("age", 25).Build()
	fmt.Printf("%#v\n\n", d)

	//  bson.D{bson.E{Key:"age", Value:bson.D{{Key:"$in", Value:[]int{18, 19, 20}}}}}
	d = query.BsonBuilder().InInt("age", 18, 19, 20).Build()
	fmt.Printf("%#v\n\n", d)

	// bson.d{bson.E{Key: "$and", Value: []any{bson.D{{Key: "x", Value: bson.D{{Key: "$ne", Value: 0}}}}, bson.D{{Key: "y", Value: bson.D{{Key: "$gt", Value: 0}}}}}}
	d = query.BsonBuilder().And(
		query.BsonBuilder().Ne("x", 0).Build(),
		query.BsonBuilder().Gt("y", 0).Build(),
	).Build()
	fmt.Printf("%#v\n\n", d)

	// bson.D{bson.E{Key:"qty", Value:bson.D{{Key:"$exists", Value:true}, bson.E{Key:"$nin", Value:[]int{5, 15}}}}}
	d = query.BsonBuilder().Exists("qty", true).NinInt("qty", 5, 15).Build()
	fmt.Printf("%#v\n\n", d)

	// elemMatch
	// bson.D{bson.E{Key: "result", Value: bson.D{bson.E{Key: "$elemMatch", Value: bson.D{bson.E{Key: "$gte", Value: 80}, bson.E{Key: "$lt", Value: 85}}}}}}
	d = query.BsonBuilder().ElemMatch("result", bsonx.D(bsonx.KV("$gte", 80), bsonx.KV("$lt", 85))).Build()
	fmt.Printf("%#v", d)
}
