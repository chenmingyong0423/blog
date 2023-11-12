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
	"github.com/chenmingyong0423/go-mongox/bsonx"
	"github.com/chenmingyong0423/go-mongox/builder/query"
	"github.com/chenmingyong0423/go-mongox/types"
)

func main() {
	query.BsonBuilder().Add(types.KV("姓名", "陈明勇")).Build() // bson.D{{Key:"姓名", Value:"陈明勇"}}

	query.BsonBuilder().And(bsonx.M("姓名", "陈明勇"), query.BsonBuilder().Gte("age", 25).Build()).Build() // bson.D{{Key:"$and", Value:[]any{bson.M{"姓名": "陈明勇"}, bson.D{{Key:"age", Value:bson.D{{Key:"$gte", Value:25}}}}}}}

	query.BsonBuilder().Or(bsonx.M("姓名", "陈明勇"), bsonx.M("姓名", "南")).Build() // bson.D{{Key:"$or", Value:[]any{bson.M{"姓名": "陈明勇"}, bson.M{"姓名": "南"}}}}

	query.BsonBuilder().InString("tags", "mongodb", "golang").Build() // bson.D{{Key:"tags", Value:bson.M{"$in": []string{"mongodb", "golang"}}}}

	query.BsonBuilder().And(query.BsonBuilder().Gt("age", 15).Build(), query.BsonBuilder().Lt("age", 25).Build()).Build() // bson.d{{Key:"$and", Value:[]any{bson.D{{Key:"age", Value:bson.M{"$gt": 15}}}, bson.D{{Key:"age", Value:bson.M{"$lt": 25}}}}}}

}
