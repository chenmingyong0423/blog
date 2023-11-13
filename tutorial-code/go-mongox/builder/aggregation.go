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
	"github.com/chenmingyong0423/go-mongox/builder/aggregation"
	"github.com/chenmingyong0423/go-mongox/builder/query"
	"github.com/chenmingyong0423/go-mongox/types"
)

func main() {

	// bson.D{bson.E{Key:"$gt", Value:[]any{"$qty", 250}}}
	gt := aggregation.BsonBuilder().Gt("$qty", 250).Build()
	fmt.Printf("%#v\n\n", gt)

	// mongo.Pipeline{bson.D{bson.E{Key:"$project", Value:bson.D{bson.E{Key:"name", Value:1}, bson.E{Key:"age", Value:1}, bson.E{Key:"qtyGt250", Value:bson.D{bson.E{Key:"$gt", Value:[]interface {}{"$qty", 250}}}}}}}}
	pipeline := aggregation.StageBsonBuilder().Project(bsonx.D(bsonx.KV("name", 1), bsonx.KV("age", 1), bsonx.KV("qtyGt250", gt))).Build()
	fmt.Printf("%#v\n\n", pipeline)

	// bson.D{bson.E{Key:"$or", Value:[]interface {}{bson.D{bson.E{Key:"score", Value:bson.D{bson.E{Key:"$gt", Value:70}, bson.E{Key:"$lt", Value:90}}}}, bson.D{bson.E{Key:"views", Value:bson.D{bson.E{Key:"$gte", Value:90}}}}}}}
	or := aggregation.BsonBuilder().Or(
		query.BsonBuilder().Gt("score", 70).Lt("score", 90).Build(),
		query.BsonBuilder().Gte("views", 90).Build(),
	).Build()
	fmt.Printf("%#v\n\n", or)

	// mongo.Pipeline{bson.D{bson.E{Key:"$match", Value:bson.D{bson.E{Key:"$or", Value:[]any{bson.D{bson.E{Key:"score", Value:bson.D{bson.E{Key:"$gt", Value:70}, bson.E{Key:"$lt", Value:90}}}}, bson.D{bson.E{Key:"views", Value:bson.D{bson.E{Key:"$gte", Value:90}}}}}}}}}, bson.D{bson.E{Key:"$group", Value:bson.D{bson.E{Key:"_id", Value:any(nil)}, bson.E{Key:"count", Value:bson.D{bson.E{Key:"$sum", Value:1}}}}}}}
	pipeline = aggregation.StageBsonBuilder().Match(or).Group(nil, bsonx.E("count", aggregation.BsonBuilder().Sum(1).Build())).Build()
	fmt.Printf("%#v\n\n", pipeline)

	// mongo.Pipeline{bson.D{bson.E{Key:"$unwind", Value:"$size"}}}
	pipeline = aggregation.StageBsonBuilder().Unwind("$size", nil).Build()
	fmt.Printf("%#v\n\n", pipeline)

	// mongo.Pipeline{bson.D{bson.E{Key:"$unwind", Value:bson.D{bson.E{Key:"path", Value:"$size"}, bson.E{Key:"includeArrayIndex", Value:"arrayIndex"}, bson.E{Key:"preserveNullAndEmptyArrays", Value:true}}}}}
	pipeline = aggregation.StageBsonBuilder().Unwind("$size", &types.UnWindOptions{
		IncludeArrayIndex:          "arrayIndex",
		PreserveNullAndEmptyArrays: true,
	}).Build()
	fmt.Printf("%#v\n\n", pipeline)
}
